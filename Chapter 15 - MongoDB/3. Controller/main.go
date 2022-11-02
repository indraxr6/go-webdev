package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Id string `json:"id"`
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := User{
		Name: "Mas Bin",
		Age: 20,
		Gender: "M",
		Id: "2",	
	}
	convert, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", convert )
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = "212"
	convert, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", convert )
}

func deleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}