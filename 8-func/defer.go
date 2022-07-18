package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer fmt.Println("1=", a+b)
	fmt.Println("2=", a*b)
	defer fmt.Println("3=", a-b)
	fmt.Println("4=", a/b)
}
