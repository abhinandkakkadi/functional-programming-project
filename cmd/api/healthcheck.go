package main

import (
	"net/http"
)

// application status, environment and version
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "available",
		"environment": app.cfg.Env,
		"version": version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		http.Error(w, "The server encountered an error and could not process the request ", http.StatusInternalServerError)
		return
	}
}