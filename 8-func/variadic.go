package main

import "fmt"

func main() {
	show(10, "20", 1, 2, 3, 4, 5)
	show(15, "hello")
	fmt.Println()
}
func show(a int, b string, v ...int) {
	fmt.Printf("%T\n", v)
	fmt.Println("%#v", v)
	fmt.Println(a)
	fmt.Println(b)
	for i, v := range v {
		fmt.Println(i, v)
	}
}
