package main

import (
	"io"
	"net/http"
)

type indomie int

func (m indomie) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "ANJENNNGGGG")
	case "/cat":
		io.WriteString(res, "INDOGMIEE")
	}

}

func main() {
	var show indomie
	http.ListenAndServe(":8080", show)
}
