# Go PostgreSQL CRUD API

This project is a simple CRUD API built with Go and PostgreSQL. It allows users to perform Create, Read, Update, and Delete operations on user data stored in a PostgreSQL database.

## Features

- Create, Read, Update, and Delete users
- RESTful API endpoints
- JSON request and response format
- PostgreSQL integration using the `pgx` driver

## Getting Started

### Prerequisites

- Go 1.16+
- PostgreSQL
- Git

### Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/your-username/go-postgres.git
    cd go-postgres
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Set up PostgreSQL:**
    - Create a PostgreSQL database and user.
    - Update the connection string in `main.go` to match your database configuration.

4. **Run the application:**
    ```sh
    go run main.go
    ```

### Project Structure





### API Endpoints

#### Create a User
- **URL:** `/user`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "name": "John Doe",
        "gender": "Male",
        "age": 30
    }
    ```
- **Response:**
    ```json
    {
        "id": 1,
        "name": "John Doe",
        "gender": "Male",
        "age": 30
    }
    ```

#### Get All Users
- **URL:** `/user/`
- **Method:** `GET`
- **Response:**
    ```json
    [
        {
            "id": 1,
            "name": "John Doe",
            "gender": "Male",
            "age": 30
        }
    ]
    ```

#### Get User by ID
- **URL:** `/user/:id`
- **Method:** `GET`
- **Response:**
    ```json
    {
        "id": 1,
        "name": "John Doe",
        "gender": "Male",
        "age": 30
    }
    ```

#### Update a User
- **URL:** `/user/:id`
- **Method:** `PUT`
- **Request Body:**
    ```json
    {
        "name": "John Smith",
        "gender": "Male",
        "age": 31
    }
    ```
- **Response:**
    ```json
    {
        "id": 1,
        "name": "John Smith",
        "gender": "Male",
        "age": 31
    }
    ```

#### Delete a User
- **URL:** `/user/:id`
- **Method:** `DELETE`
- **Response:**
    ```json
    {
        "message": "Deleted user with ID: 1"
    }
    ```

### Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
- [Go](https://golang.org/)
