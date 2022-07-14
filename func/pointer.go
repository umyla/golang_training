package main

import "fmt"

func main() {

	x := 10
	var y *int
	y = &x
	fmt.Println(&x)
	fmt.Println(y)
	*y = 20
}
