package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: We need to add unique to email and username columns
func (a *App) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("{error: bad password"))
		return
	}

	err = a.datalayer.InsertUser(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}
