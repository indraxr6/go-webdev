package main

import (
	"encoding/base64"
	"fmt"
)

func main() {	
	str := "If you're fond of sand dunes and salty air Quaint little villages here and there (You're sure) You're sure to fall in love with old Cape Cod (Old Cape Cod, that old Cape Cod)" 
	encodestd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodestd).EncodeToString([]byte(str))
	fmt.Println(len(s64))
	fmt.Println(len(str))
	fmt.Println(str)
	fmt.Println(s64)
}

