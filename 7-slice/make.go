package main

import "fmt"

func main() {
	i := make([]int, 0, 100)
	if i == nil {
		fmt.Println("i is nill")
	} else {
		fmt.Println("i is not nill")
	}
	inspectSlice("i", i)
	i = append(i, 100, 200)
	inspectSlice("i", i)
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
