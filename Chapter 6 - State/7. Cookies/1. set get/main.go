package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("./favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "lol-cookie",
		Value: "testing aja",
	})
	fmt.Fprintln(w, "Check ya browser cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	cookies, err := req.Cookie("lol-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		return
	}
	fmt.Fprintln(w, "COOKIE : ", cookies)
}