package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// Book struct to represent a book resource
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   string `json:"year"`
}

// Books slice to seed initial data
var books []Book

// getBooks function to return all books
func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// getBook function to return a single book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) // Get parameters
    for _, item := range books {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.NotFound(w, r)
}

// createBook function to add a new book
func createBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = fmt.Sprintf("%d", len(books)+1) // Generate a new ID
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

// deleteBook function to remove a book by ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(books)
}

// initializeBooks to add some initial data
func initializeBooks() {
    books = append(books, Book{ID: "1", Title: "1984", Author: "George Orwell", Year: "1949"})
    books = append(books, Book{ID: "2", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Year: "1951"})
}

func main() {
    // Initialize sample data
    initializeBooks()

    // Initialize router
    r := mux.NewRouter()

    // Route handlers
    r.HandleFunc("/api/books", getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/api/books", createBook).Methods("POST")
    r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

    // Start server
    fmt.Println("Server running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
