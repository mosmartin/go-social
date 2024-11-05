package main

import (
	"log/slog"
	"net/http"
	"time"
)

type app struct {
	config config
}

type config struct {
	addr string
}

func (a *app) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/health", a.healthCheckHandler)

	return mux
}

func (a *app) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	slog.Info("server listening on port " + a.config.addr)

	return srv.ListenAndServe()
}
