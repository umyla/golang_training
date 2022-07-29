package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

//var url = []string{"https://www.google.com/"}
var wg = &sync.WaitGroup{}

// type responce struct {
// 	url1 string
// 	resp *http.Response
// 	err  error
//}

func main() {
	//ch := make(chan string, 10)
	//respCh := make(chan responce)
	//sem := make(chan bool, 2)
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	r := []string{"https://www.google.com/", "https://go.dev/doc/effective_go", "https://www.youtube.com/"}

	for _, u := range r {
		wg.Add(4)
		go func(u string) {

			defer wg.Done()
			res, err := http.Get(u)
			if err != nil {
				log.Println(err)
				return
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
				return
			}
			defer res.Body.Close()
			fmt.Println(string(b))
			fmt.Println(res.Status)

		}(u)
		go func() {
			defer wg.Done()
			c1 <- r[0]
		}()
		go func() {
			defer wg.Done()
			c2 <- r[1]
		}()
		go func() {
			defer wg.Done()
			c3 <- r[2]
		}()
		for i := 1; i <= 3; i++ {
			select {
			case x := <-c1:

				fmt.Println(x)
			case y := <-c2:
				fmt.Println(y)
			case z := <-c3:
				fmt.Println(z)

			}
		}
	}
	wg.Wait()

}
