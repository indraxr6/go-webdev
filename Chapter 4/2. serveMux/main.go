package main

import (
	"io"
	"net/http"
)

type dog int
type cat int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ppppppppppp")
}

func (c cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "wwwww")
}

func main() {
	var d dog
	var c cat
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)
	
	http.ListenAndServe(":8080", mux)
}
