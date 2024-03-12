package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type Config struct {
	Port int
	Env  string
}

type application struct {
	cfg Config
	logger *log.Logger
}


func main() {

	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	app := &application{
		cfg: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("starting %s server on %d", cfg.Env, cfg.Port)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
