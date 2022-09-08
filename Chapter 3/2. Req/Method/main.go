package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type indomie int

var tpl *template.Template

func (m indomie) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method     string
		Submission url.Values
	}{
		r.Method,
		r.Form,
	}
	tpl.ExecuteTemplate(res, "index.gohtml", data)

}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var m indomie
	http.ListenAndServe(":6969", m)

}
