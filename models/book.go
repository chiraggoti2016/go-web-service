// models/book.go
package models

import (
    "database/sql"
    "github.com/chiraggoti2016/go-web-service/config"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   string `json:"year"`
}

// Fetch all books
func GetAllBooks() ([]Book, error) {
    rows, err := config.DB.Query("SELECT id, title, author, year FROM books")
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

// Fetch a book by ID
func GetBookByID(id int) (*Book, error) {
    var book Book
    err := config.DB.QueryRow("SELECT id, title, author, year FROM books WHERE id = ?", id).
        Scan(&book.ID, &book.Title, &book.Author, &book.Year)
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }
    return &book, nil
}

// Create a new book
func CreateBook(book Book) (int64, error) {
    res, err := config.DB.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", book.Title, book.Author, book.Year)
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}

// Update a book
func UpdateBook(id int, book Book) error {
    _, err := config.DB.Exec("UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?", book.Title, book.Author, book.Year, id)
    return err
}

// Delete a book
func DeleteBook(id int) error {
    _, err := config.DB.Exec("DELETE FROM books WHERE id = ?", id)
    return err
}
