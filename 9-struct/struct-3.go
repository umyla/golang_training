package main

import "fmt"

type student struct {
	name  string
	email string
}

func (x student) Printdata() {
	fmt.Println(x.name)
}

func main() {
	x1 := []student{
		{
			name:  "abc",
			email: "abc@email.com",
		},
		{
			name:  "xyz",
			email: "xyz@email.com",
		},
		{},
	}
	x1[2].Printdata()
	fmt.Println(x1[2])
}
