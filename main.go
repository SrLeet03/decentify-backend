package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	router := Router()

	environment := os.Getenv("ENV_VARIABLE")
	fmt.Println("Starting the Server", environment)
	err := http.ListenAndServe(":4050",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		fmt.Println(err.Error())
	}
}
