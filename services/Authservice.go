package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SrLeet03/decentify-backend/actions"
	"github.com/SrLeet03/decentify-backend/models"
	"github.com/gorilla/mux"
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
	var findid models.RegisReq
	Con.Last(&findid)
	fmt.Println(findid)
	aut.Userid = findid.Userid + 1

	// collection := client.Database("decentify").Collection("Users")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	//_, err := collection.InsertOne(ctx, aut)

	err := Con.Create(&aut)

	if err.RowsAffected == 0 || err.Error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.Error)
		return
	}
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(response).Encode(err.Error())
	// 	return
	// }
	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode("success")
}

func GetProfile(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	var aut models.AuthReq

	json.NewDecoder(request.Body).Decode(&aut)
	params := mux.Vars(request)
	token := params["token"]
	result, err := actions.ValidateToken(token)

	if result == "" || err != "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode("failed to get the profile , login first to get profile")
		return
	}
	var record models.RegisReq
	res := Con.Where("uuid = ? ", aut.Uuid).Find(&record)

	var resp models.GetProfile
	var posts []models.Posts

	res1 := Con.Where("userid = ? ", aut.Uuid).Find(&posts)

	if res != nil || res1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("failed to get the profile")
	}

	resp.Name = record.Name
	resp.Dob = record.Dob
	resp.Username = record.Username
	resp.Posts = posts

	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode(resp)
}
