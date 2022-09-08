package main

import (
	"fmt"
	"net/http"
)

type mie int

func (m mie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Text messages")
}

func main() {
	var food mie
	http.ListenAndServe("localhost:8080", food)
}
