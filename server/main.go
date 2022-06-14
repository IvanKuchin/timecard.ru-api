package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	apihandlers "github.com/ivankuchin/timecard.ru-api/api-handlers"
	configreader "github.com/ivankuchin/timecard.ru-api/config-reader"
	"go.uber.org/zap"
)

var cfg configreader.Config
var logger *zap.SugaredLogger

func SetConfig(c configreader.Config) {
	cfg = c
}

func createLogger() {
	l, _ := zap.NewDevelopment()
	logger = l.Sugar()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debugw("HTTP request: " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Run() {
	apihandlers.SetConfig(cfg)

	createLogger()
	defer logger.Sync()
	apihandlers.SetLogger(logger)

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/api/v1/login", apihandlers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/", apihandlers.DefaultHandler)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + strconv.Itoa(cfg.Listenport),
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
		Handler:      r,
	}

	c := make(chan os.Signal, 1)

	// Run our server in a goroutine so that it doesn't block.
	go func(c chan os.Signal) {
		logger.Infof("listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			logger.Info(err)
			// c <- os.Interrupt
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

	logger.Debug("shutthing down the server")
}
