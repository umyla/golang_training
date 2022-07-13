package main

import "fmt"

func main() {
	i := [5]int{5, 10, 15, 20, 25}
	fmt.Println(i[3])
	i[3] = 21    //updating values
	i[0] = 2     //updating values
	var a [3]int //fixed size
	a[0] = 1
	a[1] = 2
	fmt.Println("after updating ", i, a)
	for i, v := range a { // index,value
		fmt.Println(i, v)
	}

	for _, v := range a { // value
		fmt.Println(v)
	}

	for i := range a { // index
		fmt.Println(i)
	}
}
