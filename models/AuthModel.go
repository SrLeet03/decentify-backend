package models

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
	Name     string `json:"name" bson:"name" `
	UserID   int64  `gorm:"AUTO_INCREMENT;unique; not null"`
	Uuid     string
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
