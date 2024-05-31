package main

import (
	"log"
	"net/http"

	"my-go-projects/go-mongo/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	// creates a new router
	r := httprouter.New()

	// Initialize the UserController with a MongoDB session
	uc := controllers.NewUserController(getSession())

	//creating routes
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// Start the server
	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mgo.Session {

	// Connect to MongoDB on the default port
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	return s
}
