package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	waitForTask()
}

type responce struct {
	url  string
	resp *http.Response
	err  error
}

func waitForTask() {
	wg.Add(2)
	url := make(chan string)
	respCh := make(chan responce)

	go func() {
		defer wg.Done()
		task := <-url
		res, err := http.Get(task)

		r := responce{
			url:  task,
			resp: res,
			err:  err,
		}
		respCh <- r
	}()
	go func() {
		defer wg.Done()
		r := <-respCh
		if r.err != nil {
			log.Println(r.err)
			return
		}

		b, err := io.ReadAll(r.resp.Body)
		if err != nil {
			log.Println(r.err)
			return
		}

		defer r.resp.Body.Close()
		fmt.Println(string(b))
		fmt.Println(r.resp.Status)
	}()
	time.Sleep(2 * time.Second)
	url <- "https://www.google.com/"
	wg.Wait()
}
