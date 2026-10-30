# Golang CRUD REST API with MySQL

**Author:** [Ramesh Fadatare](https://rameshfadatare.medium.com/)\
**Published:** September 30, 2024\
**Topics:** Go, Golang, REST API, MySQL

Source article: [Golang CRUD REST API with MySQL on
Medium](https://rameshfadatare.medium.com/golang-crud-rest-api-with-mysql-36713f43f470)

## Introduction

This tutorial explains how to build a CRUD (Create, Read, Update,
Delete) REST API using Go and MySQL. It covers setting up a Go project,
connecting to a MySQL database, implementing API endpoints, and testing
them with Postman.

Further learning: - [Learn Go (Golang)
Programming](https://www.rameshfadatare.com/learn-go-programming/) -
[100+ Golang Source Code
Examples](https://www.sourcecodeexamples.net/p/golang-source-code-examples.html)

## Prerequisites

Install the following tools before starting:

1.  **Go (Golang):** Install a current version of Go.
2.  **MySQL:** Ensure the MySQL server is installed, running, and
    accessible.
3.  **Postman:** Use it to test the API endpoints.

## Step 1: Set Up the Go Project

Create a project directory and initialize the Go module:

``` bash
mkdir go-crud-api
cd go-crud-api
go mod init go-crud-api
```

This creates a new Go module for dependency management.

## Step 2: Install Required Packages

Install the MySQL driver and Gorilla Mux router:

``` bash
go get -u github.com/go-sql-driver/mysql
go get -u github.com/gorilla/mux
```

-   `github.com/go-sql-driver/mysql`: MySQL driver for Go.
-   `github.com/gorilla/mux`: HTTP routing library.

## Step 3: Set Up the MySQL Database

### 3.1 Create a Database

``` sql
CREATE DATABASE go_crud_api;
```

### 3.2 Create the `users` Table

``` sql
USE go_crud_api;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

The table stores each user's ID, name, email address, and creation
timestamp.

## Step 4: Create the Main Application File

Create a file named `main.go`:

``` go
package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

// User represents the user model for CRUD operations.
type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    CreatedAt string `json:"created_at"`
}

var db *sql.DB

func main() {
    // Initialize the database connection.
    var err error
    db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/go_crud_api")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize the router.
    router := mux.NewRouter()

    // Define API routes.
    router.HandleFunc("/users", getUsers).Methods("GET")
    router.HandleFunc("/user/{id}", getUser).Methods("GET")
    router.HandleFunc("/user", createUser).Methods("POST")
    router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
    router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

    // Start the server on port 8000.
    log.Fatal(http.ListenAndServe(":8000", router))
}
```

### Explanation

1.  **MySQL connection:** `sql.Open` initializes a database handle using
    the MySQL driver.
2.  **Routing:** Gorilla Mux maps HTTP methods and paths to handler
    functions.
3.  **Server setup:** The HTTP server listens on port `8000`.

> Replace `user:password` with your MySQL credentials. For real
> applications, use environment variables rather than hardcoding
> credentials.

## Step 5: Implement CRUD Operations

Add the following handler functions to `main.go`, below the `main`
function.

### 5.1 Fetch All Users

``` go
func getUsers(w http.ResponseWriter, r *http.Request) {
    var users []User

    rows, err := db.Query("SELECT id, name, email, created_at FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    json.NewEncoder(w).Encode(users)
}
```

### 5.2 Fetch a Single User by ID

This handler returns one user by ID. If the user does not exist, it
returns a 404 response.

``` go
func getUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var user User
    err := db.QueryRow(
        "SELECT id, name, email, created_at FROM users WHERE id = ?",
        id,
    ).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            http.NotFound(w, r)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(user)
}
```

### 5.3 Create a New User

This handler decodes the JSON request body, inserts the user into the
database, and returns the created user.

``` go
func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := db.Exec(
        "INSERT INTO users (name, email) VALUES (?, ?)",
        user.Name,
        user.Email,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, err := result.LastInsertId()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    user.ID = int(id)
    user.CreatedAt = "now" // Placeholder from the original tutorial.
    json.NewEncoder(w).Encode(user)
}
```

### 5.4 Update an Existing User

This handler updates a user's name and email by ID, then returns a JSON
response.

``` go
func updateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = db.Exec(
        "UPDATE users SET name = ?, email = ? WHERE id = ?",
        user.Name,
        user.Email,
        id,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    user.ID = 0 // The original tutorial converts the path ID to int; validate it before doing so.
    user.CreatedAt = "now" // Placeholder from the original tutorial.
    json.NewEncoder(w).Encode(user)
}
```

### 5.5 Delete a User

This handler deletes a user by ID and returns `204 No Content` when the
database operation succeeds.

``` go
func deleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
```

## Step 6: Test the CRUD API with Postman

Start the Go application, then use Postman to send the following
requests.

### Get All Users

-   **Method:** `GET`
-   **URL:** `http://localhost:8000/users`

### Get a Single User

-   **Method:** `GET`
-   **URL:** `http://localhost:8000/user/{id}`

Replace `{id}` with an existing user ID, for example `1`.

### Create a New User

-   **Method:** `POST`
-   **URL:** `http://localhost:8000/user`
-   **Body type:** `raw` → `JSON`

``` json
{
  "name": "Ramesh Fadatare",
  "email": "ramesh.fadatare@example.com"
}
```

### Update a User

-   **Method:** `PUT`
-   **URL:** `http://localhost:8000/user/{id}`
-   **Body type:** `raw` → `JSON`

``` json
{
  "name": "Ramesh Fadatare",
  "email": "ramesh@example.com"
}
```

### Delete a User

-   **Method:** `DELETE`
-   **URL:** `http://localhost:8000/user/{id}`

Replace `{id}` with the ID of the user to delete.

## Best Practices

1.  **Error handling:** Handle errors carefully and return appropriate
    HTTP status codes.
2.  **Input validation:** Validate user input before interacting with
    the database.
3.  **Logging:** Add logging to help monitor the application and
    diagnose issues.
4.  **Environment variables:** Store database credentials in environment
    variables rather than source code.

## Conclusion

This tutorial demonstrates the basic structure of a CRUD REST API using
Go and MySQL. It initializes a Go project, creates a database and table,
defines endpoints, and tests those endpoints with Postman. Possible next
steps include adding authentication, input validation, and more advanced
queries.

## References

-   [Original tutorial: Golang CRUD REST API with
    MySQL](https://www.sourcecodeexamples.net/2024/05/golang-crud-rest-api-with-mysql.html)
-   [Learn Go Programming by Ramesh
    Fadatare](https://www.rameshfadatare.com/learn-go-programming/)
-   [Golang Source Code
    Examples](https://www.sourcecodeexamples.net/p/golang-source-code-examples.html)

------------------------------------------------------------------------

**Important:** This is a cleaned-up Markdown transcription of the
supplied article, not a production-ready API. The example still needs
stronger validation, database connection checks, correct update
handling, and consistent JSON response headers before production use.
