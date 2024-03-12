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
}