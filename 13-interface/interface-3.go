package main

import (
	"fmt"
	"log"
)

type user struct {
	name  string
	email string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Println("send notifications to ", u.name, u.email, string(p))
	return len(p), nil
}

func main() {
	u := user{
		name:  "abc",
		email: "abc@email.com",
	}
	l := log.New(u, " Sales-App", log.Ltime)
	l.Println("hello")
}
