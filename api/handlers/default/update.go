package handlers_default

import (
	"CRUD-service/db"
	"net/http"
)

func UpdateDataHandler(db *db.DefaultDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is update endpoint."))
	}
}
