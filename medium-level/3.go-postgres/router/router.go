package router

import (
	"my-go-projects/go-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	// Create a new mux router
	router := mux.NewRouter()

	// Define route handlers for various endpoints
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	// Return the configured router
	return router
}
