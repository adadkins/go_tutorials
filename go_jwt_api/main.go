package main

import (
	"fmt"
	"go_tutorials/go_jwt_api/auth"
	"go_tutorials/go_jwt_api/data_layer"
	"go_tutorials/go_jwt_api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type User = models.User

type App struct {
	datalayer data_layer.DataLayer
	Router    *mux.Router
}

func main() {
	a := App{}
	dl, err := data_layer.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dl.CloseDB()

	a.datalayer = &dl

	a.RegisterRoutes()
	err = http.ListenAndServe(":8080", a.Router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}

func (a *App) RegisterRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/generate_token", a.GenerateToken).Methods("POST")
	r.HandleFunc("/register", a.RegisterUser).Methods("POST")

	userSubRouter := r.PathPrefix("/user").Subrouter()
	userSubRouter.Use(auth.Auth)
	//userSubRouter.HandleFunc("/", a.GetUser).Methods("GET")
	userSubRouter.HandleFunc("/", a.RegisterUser).Methods("PUT")

	a.Router = r
}
