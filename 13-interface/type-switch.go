package main

import (
	"fmt"
)

type student struct {
	name   string
	points float64
}

func main() {
	var user struct{}
	var s student
	checkType("hello")
	checkType(true)
	checkType(88.99)
	checkType(100)
	checkType(user)
	checkType(s)
}
func checkType(i interface{}) {
	switch v := i.(type) {
	case student:
		fmt.Println("it is a struct", v.name, v.points)
	default:
		fmt.Println("i don't know what is the type here")
	case string:
		fmt.Println(v, "is a string")
	case int:
		fmt.Println(v, " is a integer")
	case float64:
		fmt.Println(v, "is a float")
	case bool:
		fmt.Println(v, " is a bool")
	}
}
