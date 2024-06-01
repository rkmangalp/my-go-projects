package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"my-go-projects/go-mongo/models"

	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	// MongoDB client
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {

	// Create a new UserController with the given session
	return &UserController{client}
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Get the MongoDB collection
	collection := uc.client.Database("mongo-golang").Collection("users")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // Ensure the cancel function is called to release resources

	// Find all documents in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Return 500 if an error occurs
		fmt.Println("Error finding users:", err)      // Print the error
		return
	}
	defer cursor.Close(ctx) // Close the cursor when done

	// Create a slice to store the users
	var users []models.User

	// Iterate over the cursor and decode each document into a User struct
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError) // Return 500 if decoding fails
			fmt.Println("Error decoding user:", err)      // Print the error
			return
		}
		users = append(users, user)
	}

	// Check for errors during cursor iteration
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)   // Return 500 if an error occurs
		fmt.Println("Error iterating over users:", err) // Print the error
		return
	}

	// Marshal the users slice to JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)       // Return 500 if marshaling fails
		fmt.Println("Error marshaling users to JSON:", err) // Print the error
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Set the response status to 200 OK
	w.Write(usersJSON)           // Write the JSON response
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id") // Get the user ID from the route parameters

	oid, err := primitive.ObjectIDFromHex(id) // Convert the ID to a BSON ObjectID
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // Return 404 if ID is not valid
		return
	}

	u := models.User{} // Create a User model instance

	collection := uc.client.Database("mongo-golang").Collection("users")     // Get the users collection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Set a timeout of 30 seconds
	defer cancel()                                                           // Ensure the cancel function is called to release resources

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&u) // Find the user by ID in the database
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // Return 404 if user is not found
		fmt.Println(err)                   // Print the error
		return
	}

	uj, err := json.Marshal(u) // Marshal the user to JSON
	if err != nil {
		fmt.Println(err) // Print the error if marshalling fails
	}

	w.Header().Set("Content-Type", "application/json") // Set the response content type to JSON
	w.WriteHeader(http.StatusOK)                       // Set the response status to 200 OK
	fmt.Fprintf(w, "%s\n", uj)                         // Write the JSON response
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{} // Create a User model instance

	json.NewDecoder(r.Body).Decode(&u) // Decode the request body into the user model

	u.Id = primitive.NewObjectID() // Generate a new BSON ObjectID for the user

	collection := uc.client.Database("mongo-golang").Collection("users")     // Get the users collection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Set a timeout of 30 seconds
	defer cancel()                                                           // Ensure the cancel function is called to release resources

	_, err := collection.InsertOne(ctx, u) // Insert the new user into the database
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // Return 404 if insertion fails
		return
	}

	uj, err := json.Marshal(u) // Marshal the user to JSON
	if err != nil {
		fmt.Println(err) // Print the error if marshalling fails
	}

	w.Header().Set("Content-Type", "application/json") // Set the response content type to JSON
	w.WriteHeader(http.StatusCreated)                  // Set the response status to 201 Created
	fmt.Fprintf(w, "%s\n", uj)                         // Write the JSON response
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Get the user ID from the route parameters

	oid, err := primitive.ObjectIDFromHex(id) // Convert the ID to a BSON ObjectID
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // Return 404 if ID is not valid
		return
	}

	collection := uc.client.Database("mongo-golang").Collection("users")     // Get the users collection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Set a timeout of 30 seconds
	defer cancel()                                                           // Ensure the cancel function is called to release resources

	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid}) // Delete the user by ID from the database
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // Return 404 if deletion fails
		return
	}

	w.WriteHeader(http.StatusOK)                    // Set the response status to 200 OK
	fmt.Fprintf(w, "Deleted user: %s\n", oid.Hex()) // Write the response with the deleted user ID
}
