package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	var srvMux = http.NewServeMux()
	var server = http.Server{
		Addr:           ":8080",
		Handler:        srvMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	srvMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	log.Fatal(server.ListenAndServe())
}
