package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	fmt.Println("1")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("2")

	err = a.datalayer.UpdateUser(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("4")
	json.NewEncoder(w).Encode(user)
}
