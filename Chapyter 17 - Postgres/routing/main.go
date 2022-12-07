package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Book struct {
	isbn string
	title string
	author string
	price float32
}

func init() {
	var err error
	db, err := sql.Open("postgres", "postgres://indra:forces1244@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("connected to postgreszzz")
}

func main() {
	http.HandleFunc("/books", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("select * from books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.isbn, &book.title, &book.author, &book.price)
		if err != nil {	
			panic(err)
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500) 
		return
	}

	for _, book := range books {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", book.isbn, book.title, book.author, book.price)
	}
}