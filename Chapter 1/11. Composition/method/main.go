package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 9
}
func (p person) AgeDbl() int {
	return p.Age * 2
}
func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	p1 := person{
		Name: "Indra W.",
		Age:  28,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

}
