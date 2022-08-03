package handlers

import (
	"log"
	"net/http"
	"service/middleware"
	"service/web"

	"github.com/gorilla/mux"
)

func API(log *log.Logger) http.Handler {
	app := web.App{Router: mux.NewRouter()}
	m := middleware.Mid{Log: log}
	app.HandleFunc(http.MethodGet, "/check", m.Logger(check))
	return app
}
