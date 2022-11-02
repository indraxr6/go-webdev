package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "userID", 69)
	ctx = context.WithValue(ctx, "name", "bin")

	result := dbAccess(ctx)
	fmt.Fprintln(w, result)
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

func dbAccess(ctx context.Context) string {
	// uid := ctx.Value("userID").(int)
	name := ctx.Value("name").(string)
	return name
}

