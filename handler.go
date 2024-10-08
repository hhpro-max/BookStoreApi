package main

import (
	"encoding/json"
	_ "log"
	"net/http"
	"strconv"

	_ "github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `json:"id" db:"id"`
	Title  string  `json:"title" db:"title"`
	Author string  `json:"author" db:"author"`
	Price  float64 `json:"price" db:"price"`
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var book Book
	err := db.Get(&book, "SELECT * FROM books WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	_, err := db.NamedExec(`INSERT INTO books (title, author, price) 
	VALUES (:title, :author, :price)`, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
