package main

import "fmt"

type class struct {
	name  string
	age   int
	marks []int //can add slice to a struct
}

func main() {
	name1 := class{
		name:  "aaaa",
		age:   15,
		marks: []int{10, 20, 30, 40},
	}
	name1.name = "other name"
	name1.marks = append(name1.marks, 60, 70) //we can add values to slice by using append
	fmt.Println(name1)

	name2 := class{
		name:  "bbbb",
		age:   20,
		marks: []int{5, 15, 25},
	}
	name2.marks = []int{0, 1, 2, 3}
	fmt.Println(name2.name, name2.marks)

	//var name3 class
	name4 := class{} //set to default value
	fmt.Println(name4)
}
