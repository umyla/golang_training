package main

import "fmt"

type group struct {
	name string
	age  int
}

func main() {
	a := group{
		name: "ajay",
		age:  10,
	}
	b := &a //var a *group = &a address of a

	a.name = "xyz"
	b.age = "some" // we don't need * operator to access the field
	fmt.Println(b)
	fmt.Println(a)

}
