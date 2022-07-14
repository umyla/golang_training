package main

import "fmt"

type sum func(a, b int)

func main() {
	i := func(a, b int) {
		fmt.Println(a + b)
	}
	i(2, 3)
	calc(i, "hello")
	calc2(i, "other string")

}

func calc(total sum, a string) {
	total(10, 15)
	fmt.Println(a)
}

func calc2(sum func(a, b int), a string) {
	sum(30, 20)
	fmt.Println(a)
}
