package handlers_default

import (
	"CRUD-service/db"
	"fmt"
	"net/http"
)

func WriteDataHandler(defaultDB *db.DefaultDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buffer := make([]byte, r.ContentLength)
		r.Body.Read(buffer)

		fmt.Println("Received data to write:", string(buffer))

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("This is the write endpoint."))
	}
}
