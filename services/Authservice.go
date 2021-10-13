package services

import (
	"encoding/json"
	"net/http"
)

type auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var aut auth

	json.NewDecoder(request.Body).Decode(&aut)
	res := "heuuuyy"

	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode(res)

}
