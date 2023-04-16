package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	// app.infoLog.Printf("Hit the health check route!")
	fmt.Fprintf(w, "Health is good!")
}
