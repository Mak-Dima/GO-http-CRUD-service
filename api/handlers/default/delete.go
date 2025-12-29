package handlers_default

import (
	"CRUD-service/db"
	"net/http"
)

func DeleteDataHandler(db *db.DefaultDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is delete endpoint."))
	}
}
