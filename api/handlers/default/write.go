package handlers_default

import (
	"CRUD-service/db"
	"CRUD-service/pkg/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteDataHandler(defaultDB *db.DefaultDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buffer := make([]byte, r.ContentLength)
		r.Body.Read(buffer)

		fmt.Println("Received data to write:", string(buffer))
		var entity entities.DefaultEntity
		err := json.Unmarshal(buffer, &entity)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid JSON format"))
			return
		}

		err = defaultDB.WriteNewEntity(entity)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to write data to database"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Writing data to default database...\n"))
	}
}
