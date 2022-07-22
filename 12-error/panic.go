package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	total()
	add(10)
	read("This is from read func")
	fmt.Println("main func end")

}
func read(data string) {
	fmt.Println(data)
}
func add(sum int) {
	defer recoverfunc()
	a := []int{10, 20, 30}
	fmt.Println(a[10])
}
func total() {
	//defer recoverfunc()
	a := 10
	b := 15
	fmt.Println(a + b)
}

func recoverfunc() {
	r := recover()
	if r != nil {
		fmt.Println("panic happend and recoverd")
		fmt.Println(string(debug.Stack()))
	}
}
