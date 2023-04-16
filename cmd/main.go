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

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Declare instance of config struct
	var cfg config

	// Read flags into config struct
	flag.IntVar(&cfg.port, "port", 4000, "The port on which the server will listen")
	flag.StringVar(&cfg.env, "env", "development", "The environment (development|staging|production)")
	flag.Parse()

	// Declare instance of basic logger to stdout
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare instance of application struct with a config and logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare instance of new servemux
	mux := http.NewServeMux()

	// Health check route
	mux.HandleFunc("/api/v1", app.healthCheck)

	// Declare instance of http serve with port, mux, and timeout defaults
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the http server
	logger.Printf("Starting %s server on port %d", cfg.env, cfg.port)
	err := server.ListenAndServe()
	logger.Fatal(err)
}
