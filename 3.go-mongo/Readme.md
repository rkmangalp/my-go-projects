# Go MongoDB CRUD API

This project is a simple CRUD API built with Go, MongoDB, and httprouter. It allows users to perform Create, Read, Update, and Delete operations on user data stored in a MongoDB database.

## Features

- Create, Read, Update, and Delete users
- RESTful API endpoints
- JSON request and response format
- MongoDB integration using the official Go MongoDB driver

## Getting Started

### Prerequisites

- Go 1.16+
- MongoDB
- Git

### Installation

1. **Clone the repository:**
    ```sh
    git clone hhttps://github.com/rkmangalp/my-go-projects.git
    cd go-mongo
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Run the application:**
    ```sh
    go run main.go
    ```

### Project Structure

Here's a sample README file for your Go MongoDB CRUD project named go-mongo:

markdown
Copy code
# Go MongoDB CRUD API

This project is a simple CRUD API built with Go, MongoDB, and httprouter. It allows users to perform Create, Read, Update, and Delete operations on user data stored in a MongoDB database.

## Features

- Create, Read, Update, and Delete users
- RESTful API endpoints
- JSON request and response format
- MongoDB integration using the official Go MongoDB driver

## Getting Started

### Prerequisites

- Go 1.16+
- MongoDB
- Git

### Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/your-username/go-mongo.git
    cd go-mongo
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Run the application:**
    ```sh
    go run main.go
    ```

### Project Structure

├── cmd\main
| └── main.go
├── controllers
│ └── controller.go
├── models
│ └── user.go
├── go.mod
└── go.sum

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
        "id": "60d5f87e1d25b8b3f5d8b9e1",
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
            "id": "60d5f87e1d25b8b3f5d8b9e1",
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
        "id": "60d5f87e1d25b8b3f5d8b9e1",
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
        "id": "60d5f87e1d25b8b3f5d8b9e1",
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
        "message": "Deleted user: 60d5f87e1d25b8b3f5d8b9e1"
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

- [httprouter](https://github.com/julienschmidt/httprouter)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [Go](https://golang.org/)

