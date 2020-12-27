package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Author model (Struct)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Book models (Struct)
type Book struct {
	ID       string  `json:"id"`
	Serialno string  `json:"serialno"`
	Title    string  `json:"title"`
	Author   *Author `json:"author"`
}

// Initialize books (slice)
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params (id of the book which you pass in the url)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(r.Body)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // dummy id - not safe for production
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// this method is applied because we don't have a database
	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] // dummy id - not safe for production
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books) // it will return all the other books except the deleted one.
}

func main() {
	//Initialize router
	r := mux.NewRouter()

	// Dummy data of books as we don't have a database (slice)
	books = append(books, Book{ID: "124134", Serialno: "635622", Title: "Book one",
		Author: &Author{Firstname: "Shubham", Lastname: "Waje"}})
	books = append(books, Book{ID: "265637", Serialno: "767543", Title: "Book two",
		Author: &Author{Firstname: "Yash", Lastname: "Vyas"}})
	books = append(books, Book{ID: "332415", Serialno: "535625", Title: "Book three",
		Author: &Author{Firstname: "Abhay", Lastname: "Dubey"}})

	// Handel the routes
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/create", createBooks).Methods("POST")
	r.HandleFunc("/api/books/update/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/delete/{id}", deleteBooks).Methods("DELETE")

	// Initialize a server
	log.Fatal(http.ListenAndServe(":8000", r))

}
