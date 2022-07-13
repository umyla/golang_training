package main

import "fmt"

func main() {
	var x []int
	var y []int

	x = append(x, 1, 2, 3, 4, 5)
	y = append(y, 5, 4, 3, 2, 1)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
	fmt.Printf("before append %p\n", x)
	y = append(y, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println(len(y), cap(y))
	fmt.Printf("before append %p\n", y)

}
