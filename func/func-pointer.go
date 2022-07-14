package main

import "fmt"

func main() {
	x := 25
	update(&x)
	fmt.Println(x)
	update1(&x)
	fmt.Println(x)
}
func update(i *int) {
	*i = 10
	fmt.Println("after update", i)
}
func update1(i *int) {
	*i = 50
	fmt.Println("after another update", i)
}
