package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service/handlers"
	"time"
)

func main() {

	l := log.New(os.Stdout, "user: ", log.LstdFlags)
	err := startApp(l)
	if err != nil {
		panic(err)

	}
}
func startApp(log *log.Logger) error {
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		Handler:      handlers.API(log),
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
