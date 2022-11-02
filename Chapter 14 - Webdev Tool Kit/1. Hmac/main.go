package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"fmt"
)

func main() {
	code := getCode("indra@gmail.com")
	fmt.Println(code)
	code = getCode("indrx@gmail.com")
	fmt.Println(code)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("secretkeyhere"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}