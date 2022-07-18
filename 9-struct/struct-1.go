package main

import "fmt"

type emp struct {
	name    string
	id      int
	company string
}

func main() {
	emp1 := emp{
		name:    "name1",
		id:      1234,
		company: "comp1",
	}
	emp2 := emp{
		name:    "name2",
		id:      5678,
		company: "comp2",
	}
	fmt.Println(emp1.company)
	fmt.Println(emp2.name)
}
