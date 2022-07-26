package main

import "fmt"

func main() {

	var i interface{} = 100
	i = "hello string"
	i = true
	i = 99.99
	fmt.Printf("%T\n", i)
	s, ok := i.(float64)
	if !ok {
		fmt.Println("i is not s string type")
		return
	}
	fmt.Println(s)

	acceptAnything("welcome", true, 100, 99.9)
}
func acceptAnything(a ...interface{}) {
	fmt.Printf("%#v\n", a)
	for _, v := range a {
		fmt.Printf("%T\n", v)
	}
}
