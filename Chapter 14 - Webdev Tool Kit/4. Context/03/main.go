package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
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

	result, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(w, result)
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) 
	defer cancel()

	ch := make(chan int)

	go func() {
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second)

		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	case <- ctx.Done():
		return 0, ctx.Err()
	case i := <- ch:
		return i, nil
	}
}

