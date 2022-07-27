package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	c := make(chan int)
	wg.Add(3)
	go sum(10, 20, c)
	go sub(100, 10, c)
	go div(10, 5, c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z)
	wg.Wait()
}
func sum(a, b int, c chan int) {
	defer wg.Done()
	x := a + b
	c <- x
}
func sub(a, b int, c chan int) {
	defer wg.Done()
	y := a - b
	c <- y
}
func div(a, b int, c chan int) {
	defer wg.Done()
	z := a * b
	c <- z
}
