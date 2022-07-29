package main

import (
	"fmt"
	"time"
)

func main() {

}
func waitForResult() {
	c := make(chan string)
	go func() {
		time.Sleep(10 * time.Second)
		c <- "result"
	}()
	r := <-c
	fmt.Println(r)
}
