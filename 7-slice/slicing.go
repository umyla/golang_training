package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	b := a[1:5]
	fmt.Println(a, b)
	//fmt.Printf("address for a %p\n", a)
	//fmt.Printf("address for b %p\n", b)
	//a = append(a, 60, 70)
	//fmt.Println(len(a), len(b))
	//fmt.Println(a, b)
	//fmt.Printf("address for a %p\n", a)
	//fmt.Printf("address for b %p\n", b)
	//b = append(b, 1, 2, 3, 4)
	b[3] = 25
	fmt.Println(b)
	fmt.Println(a)
	//fmt.Println(len(b), b)
	//fmt.Printf("address for b %p\n", b)

}
