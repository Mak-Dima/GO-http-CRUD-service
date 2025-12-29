package handlers_default

import (
	"CRUD-service/db"
	"net/http"
)

func WriteDataHandler(defaultDB *db.DefaultDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("This is the write endpoint."))
	}
}
