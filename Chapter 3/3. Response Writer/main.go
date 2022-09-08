package main

import (
	"fmt"
	"net/http"
)

type indomie int

func (m indomie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Masbin", "from bin")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>RIGHT HERE<h1/>")
}

func main() {
	var m indomie
	http.ListenAndServe(":8080", m)
}
