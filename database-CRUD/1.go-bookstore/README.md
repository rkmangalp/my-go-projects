# Go Bookstore API

This project is a simple bookstore API built with Go, Gorilla Mux, and Gorm. It allows users to perform CRUD operations on books stored in a MySQL database.

## Features

- Create, Read, Update, and Delete books
- RESTful API endpoints
- JSON request and response format
- MySQL database integration using Gorm ORM

## Getting Started

### Prerequisites

- Go 1.16+
- MySQL
- Git

### Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/rkmangalp/my-go-projects.git
    cd go-bookstore
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Configure the database:**

    Create a MySQL database and update the database configuration in `pkg/config/app.go`:
    ```go
    d, err := gorm.Open("mysql", "username:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
    ```

4. **Run the application:**
    ```sh
    go run cmd/main.go
    ```

### Project Structure

