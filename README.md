# Service API with Golang (CRUD)

This project is a simple CRUD (Create, Read, Update, Delete) service API built with Golang.

## Features

*   **Create:** Add new users to the database.
*   **Read:** Retrieve a list of users or a single user.
*   **Update:** Modify existing user information.
*   **Delete:** Remove users from the database.

## Database

This service uses a MySQL database to store data. The connection is configured in `database/database.go`.

## Getting Started

### Prerequisites

*   Go (version 1.24.5 or later)
*   MySQL

### Installation

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    ```
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Set up the database connection in `database/database.go`.
4.  Run the application:
    ```bash
    go run main.go
    ```

### Build

To build the application into a binary file, run the following command:
```bash
go build -o service-api main.go
```

The API will be running at `http://localhost:8080`.