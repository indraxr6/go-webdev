package main

import (
	"io"
	"net/http"
)

type dog int
type cat int

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ppppppppppp")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "wwwww")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
