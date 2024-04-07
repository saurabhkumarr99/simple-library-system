package main

import (
	"encoding/json"
	"net/http"
)

// Book represents the model for a book
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

// Storage for storing books
var storage = make(map[string]Book)

// GetAllBooksFromStorage retrieves all books from storage
func GetAllBooksFromStorage() []Book {
	books := make([]Book, 0, len(storage))
	for _, book := range storage {
		books = append(books, book)
	}
	return books
}

// AddBookToStorage adds a new book to storage
func AddBookToStorage(book Book) {
	storage[book.ID] = book
}

// UpdateBookInStorage updates a book in storage by ID
func UpdateBookInStorage(id string, updatedBook Book) {
	storage[id] = updatedBook
}

// DeleteBookFromStorage deletes a book from storage by ID
func DeleteBookFromStorage(id string) {
	delete(storage, id)
}

func main() {
	// Register routes
	RegisterRoutes()

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}

// RegisterRoutes registers all the routes for the HTTP server.
func RegisterRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/books", GetAllBooksHandler)
	http.HandleFunc("/books/add", AddBookHandler)
	http.HandleFunc("/books/update", UpdateBookHandler)
	http.HandleFunc("/books/delete", DeleteBookHandler)
}

// HomeHandler handles the root route.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Simple Library System API"))
}

// GetAllBooksHandler handles the GET request for retrieving all books.
func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Dummy response for all books
	books := GetAllBooksFromStorage()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// AddBookHandler handles the POST request for adding a new book.
func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	AddBookToStorage(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// UpdateBookHandler handles the PUT request for updating a book.
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the book exists in storage
	_, exists := storage[updatedBook.ID]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Update the book in storage
	storage[updatedBook.ID] = updatedBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBookHandler handles the DELETE request for deleting a book.
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	DeleteBookFromStorage(id)

	w.WriteHeader(http.StatusNoContent)
}
