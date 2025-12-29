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

	srvMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>I am roooot</h1>"))
	})

	log.Fatal(server.ListenAndServe())
}
