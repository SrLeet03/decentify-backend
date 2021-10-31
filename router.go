package main

import (
	"fmt"

	"github.com/SrLeet03/decentify-backend/services"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("Initialise Router")

	router := mux.NewRouter()
	router.HandleFunc("/login", services.Login).Methods("POST")
	router.HandleFunc("/register", services.Register).Methods("POST")
	router.HandleFunc("/createpost/{token}", services.CreatePost).Methods("POST")
	router.HandleFunc("/profile/{token}", services.GetProfile).Methods("GET")
	router.HandleFunc("/getpost/{tag}", services.GetPosts).Methods("GET")
	router.HandleFunc("/validatoken/{tag}", services.ValidateTok).Methods("GET")

	return router
}
