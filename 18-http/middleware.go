package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", loginMid(errMid(home)))
	http.HandleFunc("/search", loginMid(search))
	panic(http.ListenAndServe(":8082", nil))
}
func loginMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("session started")
		defer log.Println("session ended")
		log.Println(r.Method)
		if r.Method != http.MethodGet {
			http.Error(w, "method allowed is get", http.StatusBadRequest)
			return
		}
		next(w, r)
	}
}
func errMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("error happend here")
		next(w, r)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	log.Println("in home handler")
	fmt.Fprintln(w, "welcome to home page")
}
func search(w http.ResponseWriter, r *http.Request) {
	log.Println("in search")
	fmt.Fprintln(w, "this is search page")
}
