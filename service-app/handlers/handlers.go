package handlers

import (
	"log"
	"net/http"
	"service/auth"
	"service/middleware"
	"service/web"

	"github.com/gorilla/mux"
)

func API(log *log.Logger, a *auth.Auth) http.Handler {
	app := web.App{Router: mux.NewRouter()}
	m := middleware.Mid{Log: log, A: a}
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(check)))))
	return app
}
