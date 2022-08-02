package main

import (
	"fmt"
	"sync"
)

var cab = 2
var wg = &sync.WaitGroup{}

func main() {
	m := &sync.Mutex{}
	names := []string{"a", "b", "c", "d"}
	for _, name := range names {
		wg.Add(1)
		go book(name, m)
	}
	wg.Wait()
	fmt.Println(cab)
}
func book(name string, m *sync.Mutex) {
	defer wg.Done()
	fmt.Println("welcome to cab service", name)
	fmt.Println("we have some offers for you", name)
	m.Lock()
	if cab >= 1 {
		fmt.Println("cab is available for", name)
		fmt.Println("booking confirmed", name)
		cab--
	} else {
		fmt.Println("cab is not available", name)
	}
	m.Unlock()
}
