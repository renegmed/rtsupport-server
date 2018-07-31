package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

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
		fmt.Println(err)
		return
	}

	// user := User {
	// 	Name: "anonymous",
	// }

	// err := r.Table("user").
	// 	Insert(user).
	// 	Exec(session)  // doesn't return response

	// response, err := r.Table("user").
	// 	Insert(user).
	// 	RunWrite(session)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// user := User {
	// 	Name: "James Moore",
	// }

	// response, _ := r.Table("user").
	// 	Get("b815caa8-a741-4d5f-a1da-c7196914a182").
	// 	Update(user).
	// 	RunWrite(session)

	// response, _ := r.Table("user").
	// 	Get("b815caa8-a741-4d5f-a1da-c7196914a182").
	// 	Delete(user).
	// 	RunWrite(session)

	// fmt.Printf("%#v\n", response)

	cursor, _ := r.Table("user").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(session)

	var changeResponse r.ChangeResponse
	for cursor.Next(&changeResponse) {
		fmt.Printf("%#v\n", changeResponse)
	}
}
