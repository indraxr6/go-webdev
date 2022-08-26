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
	err := tpl.Execute(os.Stdout, "NOTHING")
	if err != nil {
		log.Fatalln(err)
	}
}
