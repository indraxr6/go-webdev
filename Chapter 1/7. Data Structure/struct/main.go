package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name, Quote string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	buddha := sage{
		Name:  "Ghandi",
		Quote: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
