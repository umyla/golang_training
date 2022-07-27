package main

import (
	"fmt"
	"sync"
)

func main() {
	go hello()
	go hello()
	//time.Sleep(2 * time.Second)
	defer fmt.Println("hey")
	fmt.Println("hello from main")
	hello()

}
func hello() {
	//time.Sleep(2 * time.Second)
	fmt.Println("hello from func")

}
func Runner(wg *sync.WaitGroup)
