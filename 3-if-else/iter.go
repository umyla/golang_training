package main

import (
	"fmt"
	"learn-go/temp"
)

func main() {
	a, b := 3, 33
	if a < b {
		fmt.Println("b value is greater", b)
	} else {
		fmt.Println("a value is greater", a)
	}
	temp.Sub()
}
