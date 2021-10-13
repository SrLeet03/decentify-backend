package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	router := Router()

	fmt.Println("Starting the Server")
	err := http.ListenAndServe(":4050",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		fmt.Println(err.Error())
	}
}
