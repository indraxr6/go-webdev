package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func MonthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": MonthDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
