# Simple Library System

## Description
This is a simple library system API built with Go. It provides CRUD (Create, Read, Update, Delete) operations for managing books.

## Installation

### Prerequisites
- Go (version 1.22.2 or higher) installed on your machine.

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

## API Documentation

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

## Author - Saurabh Kumar Rai