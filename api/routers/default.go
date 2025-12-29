package routers

import (
	"CRUD-service/db"
	"net/http"
)

func RegisterDefaultRoutes(mux *http.ServeMux, defaultDB *db.DefaultDB) {
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("<h1>I am roooot</h1>"))
		},
	)
}
