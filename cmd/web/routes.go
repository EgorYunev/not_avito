package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EgorYunev/not_avito/config"
	"github.com/EgorYunev/not_avito/internal/auth"
	"github.com/EgorYunev/not_avito/internal/models"
	"github.com/gorilla/mux"
)

func (a *Application) start() error {
	r := mux.NewRouter()
	r.HandleFunc("/", a.home).Methods("GET")
	r.HandleFunc("/user", a.getUserById).Methods("GET")
	r.HandleFunc("/user", a.createUser).Methods("POST")
	r.HandleFunc("/auth", a.authrorize).Methods("POST")
	return http.ListenAndServe(config.ServerPort, r)
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (a *Application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (a *Application) authrorize(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)

	user := &models.User{}

	if err := dec.Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		a.Logger.Error(err.Error())
		return
	}

	ok, err := a.UserService.Authorize(user.Email, user.Password)

	if !ok || err != nil {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		a.Logger.Info("Unauthorized")
		return
	} else {
		token, err := auth.GenerateJWT(user.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			a.Logger.Error(err.Error())
			return
		}

		r.Header.Set("Token", token)

	}

}

func (a *Application) getUserById(w http.ResponseWriter, r *http.Request) {

	strId := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(strId)

	user, err := a.UserService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		a.Logger.Error(err.Error())
		return
	}

	renderJSON(w, user)

}

func (a *Application) createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		a.Logger.Error(err.Error())
		return
	}

	err := a.UserService.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		a.Logger.Error(err.Error())
		return
	}

	a.Logger.Info("New user created")

}
