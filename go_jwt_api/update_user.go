package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Name == "" || user.Password == "" || user.Username == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// hash the password
	if err := user.HashPassword(user.Password); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("{error: bad password"))
		return
	}

	// update the user in the database
	err = a.datalayer.UpdateUser(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}
