package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service/auth"
	"service/data/user"
	"service/database"
	"service/handlers"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {

	l := log.New(os.Stdout, "user: ", log.LstdFlags)
	err := startApp(l)
	if err != nil {
		panic(err)

	}

}
func startApp(log *log.Logger) error {
	// =========================================================================
	//Start Database
	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}

	uDB := user.DbService{DB: db}

	log.Println("main:started:intializing authencation support")

	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		return fmt.Errorf("reading auth private key %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parsing auth private Key %w", err)
	}
	a, err := auth.NewAuth(privateKey)
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		Handler:      handlers.API(log, a, &uDB),
	}
	shutdowm := make(chan os.Signal, 1)
	signal.Notify(shutdowm, os.Interrupt)

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("main: API listenig on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)
	case sig := <-shutdowm:
		log.Printf("main: %v: Start Stutdown", sig)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		err := api.Shutdown(ctx)
		if err != nil {
			err = api.Close()
			return fmt.Errorf("could not stop server gracefully %w", err)
		}

	}
	return nil

}
