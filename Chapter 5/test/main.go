package main

import "fmt"

type handleString struct {
	name string
	age  int
}

func (newString handleString) printString {
	fmt.Println(newString)
}

func main() {
	res := handleString{"Indra", 18}
	res.printString()
}
