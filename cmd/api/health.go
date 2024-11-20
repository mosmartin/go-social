package main

import (
	"net/http"
)

func (a *app) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     a.config.env,
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}
