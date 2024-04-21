package main

import (
	"bytes"
	"database/sql"
	"log"

	// Import the PostgreSQL driver
	_ "github.com/lib/pq"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testDB *sql.DB

func setupIntegrationTest() {
	var err error
	// Connect to the test PostgreSQL database
	testDB, err = sql.Open("postgres", "postgres://postgres:admin@localhost:5432/test_LibrarySystem?sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to test database: %v", err)
	}
}

func teardownIntegrationTest() {
	// Close the test database connection
	if testDB != nil {
		testDB.Close()
	}
}

func TestIntegrationGetAllBooksFromDB(t *testing.T) {
	// Set up integration test
	setupIntegrationTest()
	defer teardownIntegrationTest()

	// Create a new HTTP request to simulate a GET request to GetAllBooksHandler
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the GetAllBooksHandler function directly, passing in the ResponseRecorder and Request
	GetAllBooksHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body into a slice of Book structs
	var books []Book
	if err := json.Unmarshal(rr.Body.Bytes(), &books); err != nil {
		t.Fatalf("Error decoding JSON response: %v", err)
	}

	// Check the length of the books slice
	expectedLength := 7
	if len(books) != expectedLength {
		t.Errorf("Handler returned unexpected number of books: got %d want %d", len(books), expectedLength)
	}

}




func TestIntegrationAddBookHandler(t *testing.T) {
	// Set up integration test
	setupIntegrationTest()
	defer teardownIntegrationTest()

	// Create a new HTTP POST request to simulate adding a new book
	reqBody := []byte(`{"title": "New Book", "author": "Author Name", "year": "2024"}`)
	req, err := http.NewRequest("POST", "/books/add", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the AddBookHandler function directly, passing in the ResponseRecorder and Request
	AddBookHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Decode the response body to get the added book
	var addedBook Book
	if err := json.Unmarshal(rr.Body.Bytes(), &addedBook); err != nil {
		t.Fatalf("Error decoding JSON response: %v", err)
	}

}

func TestIntegrationUpdateBookHandler(t *testing.T) {
	// Set up integration test
	setupIntegrationTest()
	defer teardownIntegrationTest()

	// Create a new HTTP PUT request to simulate updating an existing book
	reqBody := []byte(`{"id": "1", "title": "Updated Title", "author": "Updated Author", "year": "2025"}`)
	req, err := http.NewRequest("PUT", "/books/update", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the UpdateBookHandler function directly, passing in the ResponseRecorder and Request
	UpdateBookHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestIntegrationDeleteBookHandler(t *testing.T) {
	// Set up integration test
	setupIntegrationTest()
	defer teardownIntegrationTest()

	// Create a new HTTP DELETE request to simulate deleting an existing book
	req, err := http.NewRequest("DELETE", "/books/delete?id=1", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the DeleteBookHandler function directly, passing in the ResponseRecorder and Request
	DeleteBookHandler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

}
