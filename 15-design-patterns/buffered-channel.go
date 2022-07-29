package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	c := make(chan string, 2)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- "hello"
		c <- "welcome"
		//c <- "ok"
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	wg.Wait()
}
