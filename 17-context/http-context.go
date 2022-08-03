package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", randomData)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println(err)
		return
	}

}
func randomData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("session started")
	defer log.Println("session ended")
	select {
	case <-time.After(1 * time.Millisecond):
		fmt.Fprintln(w, "random data")

	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx.Err())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
