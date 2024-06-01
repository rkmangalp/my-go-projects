package main

import (
	"fmt"
	"log"
	"my-go-projects/go-bookstore/pkg/config"
	"my-go-projects/go-bookstore/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// initializing the database connection
	config.Connect()
	if config.GetDB() == nil {
		log.Fatalf("Database connection is nil")
	} else {
		fmt.Println("Database connection established successfully")
	}

	//create a new router
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	//start the server

	log.Fatal(http.ListenAndServe("localhost:8000", r))
	fmt.Println("Server started on port 8000")
}
