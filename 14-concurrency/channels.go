package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan int)
	a := 20
	var a1, a2 int
	a1 = 5
	a2 = 10
	sum := a1 + a2
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- a
		ch <- sum
	}()
	go func() {
		defer wg.Done()
		x := <-ch
		y := <-ch
		fmt.Println(x, y)
	}()
	wg.Wait()
}
