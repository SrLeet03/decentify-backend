package models

type PostReg struct {
	Name         string `json:"name" bson:"name" `
	UserId       string `json:"userid" bson:"userid" `
	Body         string `json:"body" bson:"body" `
	Title        string `json:"title" bson:"title"`
	Tag          string `json:"tag" bson:"tag"`
	CreationTime string `json:"creationtime" bson:"creationtime"`
	Date         string `json:"date" bson:"date"`
	Amount       int64  `json:"amount" bson:"amount"`
	Days         int64  `json:"days" bson:"days"`
}
type Posts struct {
	PostId       int64  `json:"postid" gorm:"primary_key"`
	UserId       string `json:"userid" bson:"userid"`
	Title        string `json:"title" bson:"title"`
	Tag          string `json:"tag" bson:"tag"`
	Date         string `json:"date" bson:"date"`
	CreationTime string `json:"creationtime" bson:"creationtime"`
	Amount       int64  `json:"amount" bson:"amount"`
	Days         int64  `json:"days" bson:"days"`
}
type PostBody struct {
	PostId int64  `json:"postidc"`
	Body   string `json:"body" bson:"body" `
}
type PostTag struct {
	PostTagId int64  `json:"posttagid" gorm:"primary_key"`
	PostId    int64  `json:"postid"`
	Tag       string `json:"body" bson:"body" `
}
