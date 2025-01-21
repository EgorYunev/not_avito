package main

import (
	"encoding/json"
	"fmt"
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
	r.HandleFunc("/user/id", a.getUserById).Methods("GET")
	r.HandleFunc("/user/email", a.getUserByEmail).Methods("GET")
	r.HandleFunc("/register", a.register).Methods("POST")
	r.HandleFunc("/auth", a.authrorize).Methods("POST")
	r.HandleFunc("/ad", a.createAd).Methods("POST")
	r.HandleFunc("/ad", a.deleteAd).Methods("DELETE")
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

	email, err := auth.ParseJWT(r.Header.Get("Token"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(email)

	w.Write([]byte(email))
}

func (a *Application) authrorize(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)

	user := &models.User{}

	if err := dec.Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		a.Logger.Error(err.Error())
		return
	}

	id, err := a.UserService.Authorize(user.Email, user.Password)

	if id == 0 || err != nil {
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
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(token))

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

func (a *Application) getUserByEmail(w http.ResponseWriter, r *http.Request) {

	uEmail, err := auth.ParseJWT(r.Header.Get("Token"))

	if err != nil || uEmail == "" {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	email := r.URL.Query().Get("email")

	user, err := a.UserService.GetByEmail(email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		a.Logger.Error(err.Error())
		return
	}

	renderJSON(w, user)
}

func (a *Application) register(w http.ResponseWriter, r *http.Request) {
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

func (a *Application) createAd(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")
	email, err := auth.ParseJWT(token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, err := a.UserService.GetByEmail(email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		a.Logger.Error(err.Error())
		return
	}

	ad := &models.Ad{}
	dec := json.NewDecoder(r.Body)

	if err = dec.Decode(ad); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		a.Logger.Error(err.Error())
		return
	}
	ad.UserId = user.Id

	a.AdService.CreateAd(ad)
}

func (a *Application) deleteAd(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")

	email, err := auth.ParseJWT(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Println(email)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		a.Logger.Error(err.Error())
		return
	}

	err = a.AdService.Delete(id, email)

	if err != nil {
		http.Error(w, "You don't have permission for delete ad", http.StatusForbidden)
		a.Logger.Info(err.Error())
		return
	}

}
