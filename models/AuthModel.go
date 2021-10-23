package models

import "github.com/dgrijalva/jwt-go"

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Userid string `json:"id"`
	Token  string `json:"token"`
	Error  string `json:"error"`
}
type RegisReq struct {
	Userid   int64  `json:"id" gorm:"primary_key"`
	Name     string `json:"name" bson:"name" `
	Uuid     string `json:"uuid" bson:"uuid" `
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type JwtSchema struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}
