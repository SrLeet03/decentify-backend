package services

import (
	"encoding/json"
	"net/http"

	"github.com/SrLeet03/decentify-backend/models"
)

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var aut models.Auth

	json.NewDecoder(request.Body).Decode(&aut)
	res := "heuuuyy"

	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode(res)

}
