package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"INDRA",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
