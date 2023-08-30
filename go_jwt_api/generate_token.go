package main

import (
	"encoding/json"
	"fmt"
	"go_tutorials/go_jwt_api/auth"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *App) GenerateToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Are we inside?")
	var request TokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if email exists
	user, err := a.datalayer.GetUser(request.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// check if password correct
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		http.Error(w, credentialError.Error(), http.StatusUnauthorized)
		return
	}

	// generate a jwt
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("{jwt: \"%v\"}", tokenString)))
}
