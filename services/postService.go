package services

import (
	"encoding/json"
	"net/http"

	"github.com/SrLeet03/decentify-backend/actions"
	"github.com/SrLeet03/decentify-backend/models"
	"github.com/gorilla/mux"
)

func CreatePost(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")

	var savepost models.PostReg
	params := mux.Vars(request)
	token := params["token"]
	res, err := actions.ValidateToken(token)
	if res == "" || err != "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("failed to create the post , login first to craete post")
		return
	}
	json.NewDecoder(request.Body).Decode(&savepost)

	res, err = actions.CreatePost(savepost, Con)

	if res != "success" || err != "" {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("failed to create the post")
		return
	}
	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode("new post created successfuly")
}
