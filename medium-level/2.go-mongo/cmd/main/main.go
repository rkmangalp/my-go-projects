package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"my-go-projects/go-mongo/controllers"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Initialize the MongoDB client
	client, err := getMongoClient()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	// creates a new router
	r := httprouter.New()

	// Initialize the UserController with the MongoDB client
	uc := controllers.NewUserController(client)

	//creating routes
	r.GET("/user/", uc.GetAllUsers)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Println("Server has started on port: 8080")

	// Start the server
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

// Create a new MongoDB client and connect to the server
func getMongoClient() (*mongo.Client, error) {

	// Set a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the cancel function is called to release resources

	fmt.Println("This is the context", ctx)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil) // Ping the MongoDB server to verify the connection
	if err != nil {
		return nil, err // Return error if ping fails
	}

	return client, nil // Return the MongoDB client if connection is successful
}
