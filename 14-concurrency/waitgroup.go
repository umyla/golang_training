package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, wg)
	}
	wg.Add(2)
	go work(20, wg)
	go calc(10, 10, wg)
	wg.Wait()
}

func work(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("i am working on", i)
}

func calc(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	s := add(a, b)
	fmt.Println("total", s)
	a, err := strconv.Atoi("abc")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(a)
}

func add(a, b int) int {
	return a + b
}
