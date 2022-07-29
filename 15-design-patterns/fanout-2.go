package main

import (
	"fmt"
	"strconv"
)

//var wg = sync.WaitGroup{}

func main() {
	const work = 10
	const job = 2
	ch := make(chan string, work)
	sem := make(chan bool, job)
	go func() {
		for e := 1; e <= 10; e++ {
			go func(e int) {
				sem <- true
				ch <- "url" + strconv.Itoa(e)
			}(e)
			//time.Sleep(10 * time.Second)
			<-sem
		}
	}()
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
