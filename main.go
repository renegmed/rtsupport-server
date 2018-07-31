package main

import (
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
)

type Message struct {
	Name string      `json:"name" gorethink:"id,omitempty"`
	Data interface{} `json:"data" gorethink:"name"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})
	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)

	router.Handle("channel add", addChannel)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
