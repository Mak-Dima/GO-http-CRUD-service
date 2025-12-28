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

	srvMux.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is the read endpoint."))
	})

	srvMux.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("This is the write endpoint."))
	})

	srvMux.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is update endpoint."))
	})

	srvMux.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is delete endpoint."))
	})

	log.Fatal(server.ListenAndServe())
}
