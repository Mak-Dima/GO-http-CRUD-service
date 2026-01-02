package routers

import (
	handlers_default "CRUD-service/api/handlers/default"
	"CRUD-service/db"
	"net/http"
)

func RegisterDefaultRoutes(mux *http.ServeMux, defaultDB *db.DefaultDB) {
	const basePath = "/default"
	mux.HandleFunc(basePath+"/read", handlers_default.ReadAllDataHandler(defaultDB))
	mux.HandleFunc("POST "+basePath+"/write", handlers_default.WriteDataHandler(defaultDB))
	mux.HandleFunc("PUT "+basePath+"/update/{id}", handlers_default.UpdateDataHandler(defaultDB))
	mux.HandleFunc("DELETE "+basePath+"/delete/{id}", handlers_default.DeleteDataHandler(defaultDB))
}
