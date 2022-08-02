package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/search", searchSomething)
	http.ListenAndServe("localhost:8080", nil)
	//fmt.Println("end of the main")
}
func searchSomething(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "searching for somthing")
}
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the page"))
}
