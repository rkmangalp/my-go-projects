package main

import (
	"fmt"
	"log"
	"my-go-projects/go-postgres/router"
	"net/http"
)

func main() {
	// Initialize the router
	r := router.Router()

	// Print server start message
	fmt.Println("starting server on the port: 8080")

	// Start the server and log any errors if it fails to start
	log.Fatal(http.ListenAndServe(":8080", r))
}
