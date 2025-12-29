package handlers_default

import (
	"CRUD-service/db"
	"net/http"
)

func ReadAllDataHandler(defaultDB *db.DefaultDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := defaultDB.DataToByteSlice()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to read data"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
