package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

type sage struct {
	Name, Quote string
}

type car struct {
	Manufacturer, Model string
	Doors               int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	//tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
	tpl = template.Must(template.ParseFiles("index.gohtml"))
	tpl = tpl.Funcs(fm)

}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	a := sage{
		Name:  "Buddha",
		Quote: "The belief oassaf no beliefs",
	}
	b := sage{
		Name:  "Gandhi",
		Quote: "Be the change",
	}
	c := sage{
		Name:  "Martin Luther King",
		Quote: "Hatred never ceases with hatred but with love alone is healed",
	}
	d := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}
	e := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}
	f := car{
		Manufacturer: "Tesla",
		Model:        "Type X",
		Doors:        4,
	}

	sages := []sage{a, b, c}
	cars := []car{d, e, f}

	data := items{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
