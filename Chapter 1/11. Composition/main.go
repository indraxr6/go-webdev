package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	p1 := struct {
		Name string
		Age  int
	}{
		"Indra",
		25,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

}
