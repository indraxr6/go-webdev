package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bag", bag)
	http.HandleFunc("/bagged", bagged)
	http.ListenAndServe(":8080", nil)
	
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("req method foo:",  req.Method)
}

func bag(w http.ResponseWriter, req *http.Request) {
	fmt.Print("req method bag:",  req.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func bagged(w http.ResponseWriter, req *http.Request) {
	fmt.Print("req method bagged:",  req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}