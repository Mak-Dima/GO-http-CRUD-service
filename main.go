package main

import (
	"log"
	"net/http"
	"time"

	"CRUD-service/api/routers"
	"CRUD-service/db"
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
	var defaultDB = db.NewDefaultDB()
	if defaultDB == nil {
		log.Fatal("Failed to initialize the default database")
	}

	routers.RegisterDefaultRoutes(srvMux, defaultDB)

	srvMux.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		data, err := defaultDB.DataToByteSlice()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to read data"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
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
