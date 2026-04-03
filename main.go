package main

import (
	"log"
	"net/http"
)

func handleRoot() {


}

func main() {
	const port = "8080"

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("Serving on port: %s\n", port)
	srv.ListenAndServe()
}