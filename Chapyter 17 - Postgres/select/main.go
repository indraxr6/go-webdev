package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn string
	title string
	author string
	price float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://indra:forces1244@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("connected to postgreszzz")

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
		panic(err)
	}

	for _, book := range books {
		fmt.Printf("%s, %s, %s, $%.2f\n", book.isbn, book.title, book.author, book.price)
	}
}