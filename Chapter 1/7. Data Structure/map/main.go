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
	sages := map[string]string{
		"India":    "Ghandi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Christ":   "Jesus",
		"Prophet":  "Mohammad",
	}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
