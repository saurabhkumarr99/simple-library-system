package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddBookHandler(t *testing.T) {
	// Create a new HTTP POST request with a book payload
	book := Book{Title: "Test Book", Author: "Test Author", Year: "2022"}
	bookJSON, _ := json.Marshal(book)
	req, err := http.NewRequest("POST", "/books/add", bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler instance
	handler := http.HandlerFunc(AddBookHandler)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	var responseBook Book
	err = json.Unmarshal(rr.Body.Bytes(), &responseBook)
	if err != nil {
		t.Errorf("error decoding response JSON: %v", err)
	}

	if responseBook.Title != book.Title || responseBook.Author != book.Author || responseBook.Year != book.Year {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(bookJSON))
	}
}

func TestUpdateBookHandler(t *testing.T) {
	// Create a new HTTP PUT request with an updated book payload
	updatedBook := Book{ID: "1", Title: "Updated Book", Author: "Updated Author", Year: "2023"}
	bookJSON, _ := json.Marshal(updatedBook)
	req, err := http.NewRequest("PUT", "/books/update?id=1", bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler instance
	handler := http.HandlerFunc(UpdateBookHandler)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var responseBook Book
	err = json.Unmarshal(rr.Body.Bytes(), &responseBook)
	if err != nil {
		t.Errorf("error decoding response JSON: %v", err)
	}

	if responseBook.ID != updatedBook.ID || responseBook.Title != updatedBook.Title || responseBook.Author != updatedBook.Author || responseBook.Year != updatedBook.Year {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBook, updatedBook)
	}
}

func TestDeleteBookHandler(t *testing.T) {
	// Create a new HTTP DELETE request
	req, err := http.NewRequest("DELETE", "/books/delete?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler instance
	handler := http.HandlerFunc(DeleteBookHandler)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

// MockDB simulates a database for testing purposes
type MockDB struct {
	Books []Book
}

// GetAllBooksHandler returns a list of books (mock implementation)
func (m *MockDB) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(m.Books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
