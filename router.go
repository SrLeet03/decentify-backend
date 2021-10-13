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

	return router
}
