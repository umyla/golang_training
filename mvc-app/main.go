package main

import (
	"learn/controllers"
	"learn/middlewares"
	"net/http"
)

func main() {
	http.HandleFunc("/home", middlewares.LoginMid(middlewares.ErrMid(controllers.Home)))
	http.HandleFunc("/users", controllers.GetUser)

	panic(http.ListenAndServe(":8082", nil))
}
