package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Health check route
	mux.HandleFunc("/api/v1", app.healthCheck)

	return mux
}
