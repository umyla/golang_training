package main

import (
	"learn-go/mvc-app/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/home", middleWares.LoginMid(middleWares.ErrMid(controllers.Home)))
	http.HandleFunc("/users", controllers.GetUser)

	panic(http.ListenAndServe(":8082", nil))
}
