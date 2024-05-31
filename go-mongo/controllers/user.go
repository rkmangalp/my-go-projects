package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"my-go-projects/go-mongo/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	// MongoDB session
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {

	// Create a new UserController with the given session
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Get the user ID from the route parameters
	id := p.ByName("id")

	// Check if the ID is a valid BSON ObjectID
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	// Convert the ID to a BSON ObjectId
	oid := bson.ObjectIdHex(id)

	// Create a User model instance
	u := models.User{}

	// Find the user by ID in the database
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
	}

	// Marshal the user to JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Create a User model instance
	u := models.User{}

	// Decode the request body into the user model
	json.NewDecoder(r.Body).Decode(&u)

	// Generate a new BSON ObjectId for the user
	u.Id = bson.NewObjectId()

	// Insert the new user into the database
	if err := uc.session.DB("mongo-golang").C("users").Insert(u); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Marshal the user to JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user: ", oid, "\n")

}
