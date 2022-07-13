package main

import "fmt"

func main() {
	var a []int
	b := [10]int{2, 4, 6, 8, 10, 12}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b[1:4])
	a = append(a, 11, 33, 55) //adding more values to slice
	b = append(b, 16)         //adding
	fmt.Println("after append", a, "after append", b)
}
