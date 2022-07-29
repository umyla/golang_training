package main

import (
	"fmt"
	"strconv"
)

func main() {

	const job = 10
	ch := make(chan string, job)
	for work := 1; work <= 10; work++ {
		go func(w int) {
			ch <- "work" + strconv.Itoa(w)
		}(work)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}
