package handlers

import (
	"log"
	"net/http"
	"service/auth"
	"service/data/user"
	"service/middleware"
	"service/web"

	"github.com/gorilla/mux"
)

func API(log *log.Logger, A *auth.Auth, db *user.DbService) http.Handler {
	app := web.App{Router: mux.NewRouter()}
	m := middleware.Mid{
		Log: log,
		A:   A,
	}
	uh := userHandlers{
		DbService: db,
		Auth:      A,
	}

	// func API(log *log.Logger, a *auth.Auth) http.Handler {
	// 	app := web.App{Router: mux.NewRouter()}
	// 	m := middleware.Mid{Log: log, A: a}
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(m.HasRole(check, auth.RoleAdmin, auth.RoleUser))))))
	app.HandleFunc(http.MethodPost, "/create", m.Logger(m.Error(m.Panic(uh.SignUp))))
	app.HandleFunc(http.MethodPost, "/login", m.Logger(m.Error(m.Panic(uh.Login))))
	app.HandleFunc(http.MethodGet, "/add", m.Logger(m.Error(m.Panic(m.Authenticate(uh.AddInventory)))))
	app.HandleFunc(http.MethodGet, "/view", m.Logger(m.Error(m.Panic(m.Authenticate(uh.ViewInventory)))))
	return app
}
