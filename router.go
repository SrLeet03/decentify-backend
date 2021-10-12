package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("Initialise Router")

	router := mux.NewRouter()

	router.HandleFunc("/login", services.login).Methods("POST")

	return router
}
