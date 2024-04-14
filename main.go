package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

// Book represents the model for a book
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

// DB connection
var db *sql.DB

func init() {
	var err error
	// Connect to the PostgreSQL database
	db, err = sql.Open("postgres", "postgres://postgres:admin@localhost:5432/LibrarySystem?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllBooksFromDB retrieves all books from the database
func GetAllBooksFromDB() ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// AddBookToDB adds a new book to the database
func AddBookToDB(book Book) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO books(title, author, year) VALUES($1, $2, $3) RETURNING id", book.Title, book.Author, book.Year).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateBookInDB updates a book in the database by ID
func UpdateBookInDB(id string, updatedBook Book) error {
	// Convert the ID from string to integer for database query
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4", updatedBook.Title, updatedBook.Author, updatedBook.Year, idInt)
	return err
}

// DeleteBookFromDB deletes a book from the database by ID
func DeleteBookFromDB(id int) error {
	_, err := db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}

func main() {
	// Register routes
	RegisterRoutes()

	defer db.Close()

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// RegisterRoutes registers all the routes for the HTTP server.
func RegisterRoutes() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/books", GetAllBooksHandler)
	http.HandleFunc("/books/add", AddBookHandler)
	http.HandleFunc("/books/update", UpdateBookHandler)
	http.HandleFunc("/books/delete", DeleteBookHandler)

	// User routes
	http.HandleFunc("/users", GetAllUsersHandler)
	http.HandleFunc("/users/add", AddUserHandler)
	http.HandleFunc("/users/update", UpdateUserHandler)
	http.HandleFunc("/users/delete", DeleteUserHandler)

	// Login route
	http.HandleFunc("/login", LoginHandler)
}

// HomeHandler handles the root route.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Simple Library System API"))
}

// GetAllBooksHandler handles the GET request for retrieving all books.
func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve all books from the database
	books, err := GetAllBooksFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return books as JSON response
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

	// Add book to the database
	id, err := AddBookToDB(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the generated ID to the book
	book.ID = fmt.Sprintf("%d", id)

	// Return added book as JSON response
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

	// Update the book in the database
	err = UpdateBookInDB(updatedBook.ID, updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBookHandler handles the DELETE request for deleting a book.
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// Convert ID to int
	idInt := 0
	_, err := fmt.Sscanf(id, "%d", &idInt)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Delete the book from the database
	err = DeleteBookFromDB(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// User represents the model for a user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// AddUserHandler handles the POST request for adding a new user.
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add user to the database
	id, err := AddUserToDB(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the generated ID to the user
	user.ID = id

	// Return added user as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetAllUsersHandler handles the GET request for retrieving all users.
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve all users from the database
	users, err := GetAllUsersFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return users as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateUserHandler handles the PUT request for updating a user.
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user in the database
	err = UpdateUserInDB(updatedUser.ID, updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUserHandler handles the DELETE request for deleting a user.
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// Convert ID to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Delete the user from the database
	err = DeleteUserFromDB(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllUsersFromDB retrieves all users from the database
func GetAllUsersFromDB() ([]User, error) {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// AddUserToDB adds a new user to the database
func AddUserToDB(user User) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO users(username, password) VALUES($1, $2) RETURNING id", user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateUserInDB updates a user in the database by ID
func UpdateUserInDB(id int, updatedUser User) error {
	_, err := db.Exec("UPDATE users SET username=$1, password=$2 WHERE id=$3", updatedUser.Username, updatedUser.Password, id)
	return err
}

// DeleteUserFromDB deletes a user from the database by ID
func DeleteUserFromDB(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

// AuthenticateUser checks if the provided username and password are valid.
func AuthenticateUser(username, password string) bool {

	var dbUsername, dbPassword string
	err := db.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&dbUsername, &dbPassword)
	if err != nil {
		return false // User not found or error occurred
	}
	return password == dbPassword
}

// LoginHandler handles the POST request for user login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body to extract username and password
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the user's credentials
	authenticated := AuthenticateUser(user.Username, user.Password)
	if !authenticated {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
