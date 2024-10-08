package main

import (
	"log"
	"net/http"
)

func main() {
	ConnectDB()
	mux := http.NewServeMux()

	mux.HandleFunc("/books", GetBooks)
	mux.HandleFunc("/book", GetBook)
	mux.HandleFunc("/book/create", CreateBook)
	mux.HandleFunc("/book/delete", DeleteBook)

	loggedMux := LoggingMiddleware(mux)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatal(err)
	}
}
