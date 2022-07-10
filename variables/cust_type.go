package main

import (
	"fmt"
)

type id int

func main() {

	var emp1 id = 1234
	i := 4321
	emp1 = id(i)
	fmt.Println(emp1)
	hello()
}

func hello() {
	var emp1 id
	emp1 = 9999
	fmt.Println(emp1)
}
