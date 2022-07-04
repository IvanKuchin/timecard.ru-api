package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	apihandlers "github.com/ivankuchin/timecard.ru-api/api-handlers"
	configreader "github.com/ivankuchin/timecard.ru-api/config-reader"
	"github.com/ivankuchin/timecard.ru-api/logs"
)

var cfg configreader.Config

func SetConfig(c configreader.Config) {
	cfg = c
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger.yaml")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs.Sugar.Debugw("HTTP request: " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Run() {
	apihandlers.SetConfig(cfg)

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/api/v1/login", apihandlers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/agency/sow/", apihandlers.AgencySowListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/sow/{key}", apihandlers.AgencySowListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/timecard/", apihandlers.AgencyTimecardListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/timecard/{key}", apihandlers.AgencyTimecardListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/bt/", apihandlers.AgencyBTListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/bt/{key}", apihandlers.AgencyBTListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/subcontractor_by_sow/{key}", apihandlers.AgencySubcBySowListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/info", apihandlers.AgencyInfoHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/cost_centers", apihandlers.AgencyInfoHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/invoices/service/cost_center/{key}", apihandlers.InvoiceServiceHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/invoices/service/{key}", apihandlers.InvoiceServiceDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/invoices/bt/cost_center/{key}", apihandlers.InvoiceBTHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/agency/invoices/bt/{key}", apihandlers.InvoiceBTDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/", apihandlers.DefaultHandler)

	r.HandleFunc("/api/v1/swagger.yaml", serveSwaggerFile)
	opts := middleware.SwaggerUIOpts{
		BasePath: "/api/v1/",
		SpecURL:  "swagger.yaml",
	}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/api/v1/docs", sh)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + strconv.Itoa(cfg.Listenport),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      r,
	}

	c := make(chan os.Signal, 1)

	// Run our server in a goroutine so that it doesn't block.
	go func(c chan os.Signal) {
		logs.Sugar.Debugf("listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			logs.Sugar.Info(err)
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

	logs.Sugar.Debug("shutthing down http-server")
}
