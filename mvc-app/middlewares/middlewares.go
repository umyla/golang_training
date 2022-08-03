package middlewares

import (
	"log"
	"net/http"
)

func LoginMid(next http.HandlerFunc) http.HandlerFunc {
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
func ErrMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("error happend here")
		next(w, r)
	}
}
