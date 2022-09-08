package main

import (
	"io"

	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", pic)
	http.ListenAndServe(":8080", nil)
}

func pic(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charsset=utf-8")
	io.WriteString(w, `<img src="bruhh.jpeg">`)	
}

func memePic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("bruhh.jpeg")
	if err != nil {
		http.Error(w, "file nah found", 404)
	}
	defer f.Close()
	io.Copy(w, f)

}