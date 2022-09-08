package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main2() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>` + name + `<h1>
	</body>
	</html>
	`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("failed creating file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
