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
	b := &a //var b *group = &a address of a
	fmt.Println(a, b)
	a.name = "some"
	b.age = 26 // we don't need * operator to access the field
	fmt.Println(b)
	fmt.Println(a)

}
