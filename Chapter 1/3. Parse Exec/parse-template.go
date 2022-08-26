package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	} else {
		err = tpl.Execute(os.Stdout, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
