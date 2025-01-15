package main

import (
	"net/http"

	"github.com/EgorYunev/not_avito/config"
	"github.com/gorilla/mux"
)

func (a *Application) start() error {
	r := mux.NewRouter()
	r.HandleFunc("/", a.home).Methods("GET")
	return http.ListenAndServe(config.ServerPort, r)
}

func (a *Application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
