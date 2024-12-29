package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetMovies(t *testing.T) {
	// Arrange: Seed some initial data
	movies = []Movie{
		{ID: "1", Title: "Movie One"},
		{ID: "2", Title: "Movie Two"},
	}

	req, _ := http.NewRequest("GET", "/movies", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getMovies)

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}

	var result []Movie
	_ = json.Unmarshal(rr.Body.Bytes(), &result)
	if len(result) != 2 {
		t.Errorf("expected 2 movies, got %d", len(result))
	}
}

func TestCreateMovie(t *testing.T) {
	// Arrange: Create a new movie payload
	newMovie := `{
		"id": "3",
		"isbn": "12345",
		"title": "Movie Three",
		"director": {
			"firstname": "John",
			"lastname": "Doe"
		}
	}`

	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer([]byte(newMovie)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createMovie)

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}

	var created Movie
	_ = json.Unmarshal(rr.Body.Bytes(), &created)

	// Validate the movie fields
	if created.ID != "3" || created.Title != "Movie Three" || created.Isbn != "12345" {
		t.Errorf("movie not created properly, got %+v", created)
	}

	// Validate the director fields explicitly
	if created.Director == nil {
		t.Errorf("director not set, got %+v", created)
	} else if created.Director.Firstname != "John" || created.Director.Lastname != "Doe" {
		t.Errorf("director fields not set properly, got %+v", created.Director)
	}
}

func TestDeleteMovies(t *testing.T) {
	// Arrange: Seed initial data and delete a movie
	movies = []Movie{
		{ID: "1", Title: "Movie One"},
		{ID: "2", Title: "Movie Two"},
	}
	req, _ := http.NewRequest("DELETE", "/movies/1", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", deleteMovie)
	router.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}

	var remaining []Movie
	_ = json.Unmarshal(rr.Body.Bytes(), &remaining)
	if len(remaining) != 1 || remaining[0].ID != "2" {
		t.Errorf("movie not deleted properly, got %+v", remaining)
	}
}
