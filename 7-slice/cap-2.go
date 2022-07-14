package main

import "fmt"

func main() {
	x := []int{11, 22, 33, 44, 55, 66}
	y := x[2:5:5]
	//fmt.Println(x, len(x), cap(x))
	//fmt.Println(y, len(y), cap(y))
	y = append(y, 777, 888)
	fmt.Println(y, len(y), cap(y))
	calc("x", x)
	calc("y", y)

}
func calc(name string, slice []int) {
	fmt.Printf("name %v len %d cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
