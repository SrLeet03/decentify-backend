package services

import (
	"encoding/json"
	"net/http"

	"github.com/SrLeet03/decentify-backend/actions"
	"github.com/SrLeet03/decentify-backend/models"
)

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var aut models.LoginReq

	json.NewDecoder(request.Body).Decode(&aut)
	var record models.RegisReq
	resp := Con.Where("username = ? AND password = ?", aut.Username, aut.Password).Find(&record)

	if resp.RowsAffected == 0 {
		res := models.LoginResponse{
			Error: "Can't find the user for this username and password",
		}
		response.WriteHeader(http.StatusAccepted)
		json.NewEncoder(response).Encode(res)
		return
	}

	res := actions.Login(aut)
	res.Userid = record.Uuid
	if res.Error != "" {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(res)
		return
	}
	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode(res)
}
func Register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var aut models.RegisReq

	json.NewDecoder(request.Body).Decode(&aut)

	aut.Uuid = actions.GenerateUUID()

	aut.UserID = 0
	// collection := client.Database("decentify").Collection("users")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// _, err := collection.InsertOne(ctx, aut)
	err := Con.Create(&aut)

	if err.RowsAffected == 0 || err.Error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(("failed"))
	}
	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode("success")
}
