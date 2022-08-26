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
	sages := []string{"Gandhi", "Mohammad", "Jesus", "Buddha"}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)

	}
}
