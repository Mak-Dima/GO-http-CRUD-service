package handlers_default

import (
	"CRUD-service/db"
	"net/http"
)

func DeleteDataHandler(db *db.DefaultDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		err := db.DeleteEntity(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error deleting entity."))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deleted entity with ID: " + id))
	}
}
