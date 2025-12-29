package routers

import (
	handlers_default "CRUD-service/api/handlers/default"
	"CRUD-service/db"
	"net/http"
)

func RegisterDefaultRoutes(mux *http.ServeMux, defaultDB *db.DefaultDB) {

	mux.HandleFunc("/read", handlers_default.ReadAllDataHandler(defaultDB))
	mux.HandleFunc("/write", handlers_default.WriteDataHandler(defaultDB))
	mux.HandleFunc("/update", handlers_default.UpdateDataHandler(defaultDB))
	mux.HandleFunc("/delete", handlers_default.DeleteDataHandler(defaultDB))
}
