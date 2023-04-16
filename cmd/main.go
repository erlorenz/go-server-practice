package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	addr string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":3000", "The port on which the server will listen")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health is good!")
	})

	log.Printf("Listening on port %s", cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
