package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func Double(x int) int {
	return x + x
}

func Square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl":  Double,
	"fsq":   Square,
	"fsqrt": sqRoot,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}

}
