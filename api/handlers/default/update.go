package handlers_default

import (
	"CRUD-service/db"
	"CRUD-service/pkg/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateDataHandler(db *db.DefaultDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var buffer = make([]byte, r.ContentLength)
		r.Body.Read(buffer)

		var entity entities.DefaultEntity
		err := json.Unmarshal(buffer, &entity)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		entity.ID = id

		fmt.Println(entity)

		err = db.UpdateEntity(entity)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Updated: " + entity.Name))
	}
}
