package main

import "gopkg.in/mgo.v2"

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")	
}