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
	goals := struct {
		Score1, Score2 int
	}{
		1, 6,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", goals)
	if err != nil {
		log.Fatalln(err)
	}

}
