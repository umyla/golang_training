package main

import "fmt"

func main() {
	student("jen", 25, 500)
	x := sum(10, 10)
	fmt.Println(x)
}
func student(name string, age int, marks int) {
	fmt.Println(name, age, marks)
}

func sum(a int, b int) int {
	s := a + b
	return s
}
