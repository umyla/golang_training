package main

import "fmt"

func main() {
	var slice = []string{"creating", "a", "slice"}
	fmt.Println(len(slice), cap(slice))
	fmt.Println(slice)
	slice = append(slice, "after", "first", "append")
	fmt.Println(len(slice), cap(slice))
	fmt.Println(slice)
}
