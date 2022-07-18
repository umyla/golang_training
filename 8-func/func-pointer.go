package main

import "fmt"

func main() {
	x := "hello"
	update(&x)
	fmt.Println(x)
	update1(&x)
	fmt.Println(x)
}
func update(i *string) {
	*i = "ok"
	fmt.Println("after update", i)
}
func update1(i *string) {
	*i = "hai"
	fmt.Println("after another update", i)
}
