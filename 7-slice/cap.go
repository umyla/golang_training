package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	b := a[1:4]
	a = append(a, 5, 15)
	inter("a", a)
	fmt.Println(a)
	b = append(b, 35, 65)
	fmt.Println(b)
	inter("b", b)
}
func inter(name string, slice []int) {
	fmt.Printf("name %v len %d cap %d \n", name, len(slice), cap(slice))
	println(slice)
	println()
}
