package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var url = []string{"https://www.google.com/"}
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

	for _, u := range url {
		wg.Add(1)
		go func(u string) {
			//url := <-ch
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
	}
	// go func() {
	// 	for v := range ch {
	// 		fmt.Println(v)
	// 	}

	// }()
	wg.Wait()

}
