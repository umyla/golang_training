package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	list := os.Args[1:]
	fmt.Println(list[4])
	fmt.Println(list[3])
	fmt.Println(list[1])
	fmt.Println(len(list))
}
