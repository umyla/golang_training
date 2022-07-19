package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Fullname string          `json:"first_name"`
	Password string          `json:"-"`
	Perm     map[string]bool `json:"perms,omitempty"`
}

func main() {
	emp := []user{
		{
			Fullname: "abhi",
			Password: "abcd",
			Perm:     map[string]bool{"admin": true},
		},
		{
			Fullname: "jay",
			Password: "xyz",
			Perm:     map[string]bool{"read": true},
		},
		{
			Fullname: "jack",
			Password: "dcba",
			Perm:     map[string]bool{"write": false},
		},
	}
	jsonData, err := json.MarshalIndent(emp, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(emp)
	fmt.Println(string(jsonData))
}
