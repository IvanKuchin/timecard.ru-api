package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	apihandlers "github.com/ivankuchin/timecard.ru-api/api-handlers"
	"github.com/ivankuchin/timecard.ru-api/config_reader"
)

var cfg config_reader.Config

func SetConfig(c config_reader.Config) {
	cfg = c
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("HTTP request: %v\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello back\n")
}

func Run() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/api/v1/user", UserHandler)
	r.HandleFunc("/", apihandlers.DefaultHandler)

	srv := &http.Server{
		Addr:         "127.0.0.1:" + strconv.Itoa(cfg.Listenport),
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
		Handler:      r,
	}

	c := make(chan os.Signal, 1)

	// Run our server in a goroutine so that it doesn't block.
	go func(c chan os.Signal) {
		log.Printf("listening on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("ERROR: %s\n", err)
			c <- os.Interrupt
		}
	}(c)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	log.Println("shutthing down the server")
}
