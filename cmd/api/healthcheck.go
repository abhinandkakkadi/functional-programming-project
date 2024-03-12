package main

import (
	"fmt"
	"net/http"
)

// application status, environment and version
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment %s\n", app.cfg.Env)
	fmt.Fprintf(w, "version: %s\n", version)

	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.cfg.Env, version)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}