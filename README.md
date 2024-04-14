# Simple Library System

## Description
This is a simple library system API built with Go. It provides CRUD (Create, Read, Update, Delete) operations for managing books.

## Installation

### Prerequisites
- Go (version 1.22.2 or higher) installed on your machine.
- Postgres , Pgadmin

### Steps
1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/saurabhkumarr99/simple-library-system

2. Navigate to the project directory:

   ```bash
   cd simple-library-system

3. Run Project - Open terminal 

   ```bash
   go run main.go

## SQL Queries

### Users Table
   ```sql
   -- Create users table
      CREATE TABLE IF NOT EXISTS users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(255) NOT NULL UNIQUE,
       password VARCHAR(255) NOT NULL
     );

```

### Users Table
```sql
-- Create books table
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year VARCHAR(4) NOT NULL
);

```
## API Documentation

### Books API

### 1- Get All Books
- **URL:** `http://localhost:8080/books`
- **Method:** GET
- **Description:** Retrieve all books stored in the library system.
- **Response:**
  ```json
  [
      {
          "id": "1",
          "title": "Book 1",
          "author": "Author 1",
          "year": "2022"
      },
      {
          "id": "2",
          "title": "Book 2",
          "author": "Author 2",
          "year": "2023"
      },
      
  ]

### 2- Get All Books
- **URL:** `http://localhost:8080/books`
- **Method:** GET
- **Description:** Retrieve all books stored in the library system.
- **Response:**
  ```json
  [
      {
          "id": "1",
          "title": "Book 1",
          "author": "Author 1",
          "year": "2022"
      },
      {
          "id": "2",
          "title": "Book 2",
          "author": "Author 2",
          "year": "2023"
      },
      
  ]

### 3- Update a Book
- **URL:** `http://localhost:8080/books/update`
- **Method:** PUT
- **Description:** Update an existing book in the library system.
- **Request Body:**
  ```json
  {
      "id": "1",
      "title": "Updated Book",
      "author": "Updated Author",
      "year": "2022"
  }

### 4- Delete a Book
- **URL:** `/books/delete?id={book_id}`
- **Method:** DELETE
- **Description:** Delete a book from the library system by its ID.
- **Response Status:** 204 No Content

### Users API

### 1- Get All Users
- **URL:** `http://localhost:8080/users`
- **Method:** GET
- **Description:** Retrieve all users stored in the system.
- **Response:**
  ```json
  [
      {
          "id": 1,
          "username": "user1",
          "password": "password123"
      },
      {
          "id": 2,
          "username": "user2",
          "password": "password456"
      },
      
  ]

### 2- Add User
- **URL:** `http://localhost:8080/users/add`
- **Method:** POST
- **Description:** Add a new user to the system.
- **Request Body:**
  ```json
  {
      "username": "newuser",
      "password": "newpassword"
  }

### 3- Update User
- **URL:** `http://localhost:8080/users/update`
- **Method:** PUT
- **Description:** Update an existing user in the system.
- **Request Body:**
  ```json
  {
      "id": 1,
      "username": "updateduser",
      "password": "updatedpassword"
  }

### 4- Delete User
- **URL:** `http://localhost:8080/users/delete?id=1`
- **Method:** DELETE
- **Description:** Delete a user from the system by its ID.
- **Response Status:** 204 No Content

### Login API

### 5- User Authentication and Login
- **URL:** `http://localhost:8080/login`
- **Method:** POST
- **Description:** Authenticate user credentials and login.
- **Request Body:**
  ```json
  {
      "username": "example",
      "password": "password123"
  }

## Author - Saurabh Kumar Rai