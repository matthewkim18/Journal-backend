package main

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo"
)

var db *mgo.Database

func main() {
	url := "localhost:27017"
	session, err := mgo.Dial(url)

	if err != nil {
		log.Fatal(err)
	}
	db = session.DB("journalDB")

	router := NewRouter(AllRoutes())

	log.Fatal(http.ListenAndServe(":8080", router))
}
