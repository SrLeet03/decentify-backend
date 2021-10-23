package models

type PostReg struct {
	Name         string   `json:"name" bson:"name" `
	UserId       string   `json:"postid" bson:"userid" `
	Body         string   `json:"body" bson:"body" `
	Tile         string   `json:"tile" bson:"tile"`
	Tags         []string `json:"tags" bson:"tags"`
	CreationTime string   `json:"creationtime" bson:"creationtime"`
	Amount       int64    `json:"amount" bson:"amount"`
	Days         int64    `json:"days" bson:"days"`
}
type Posts struct {
	PostId       int64    `json:"postid" gorm:"primary_key"`
	Name         string   `json:"name" bson:"name" `
	UserId       string   `json:"userid" bson:"userid" `
	Tile         string   `json:"tile" bson:"tile"`
	Tags         []string `json:"tags" bson:"tags"`
	CreationTime string   `json:"creationtime" bson:"creationtime"`
	Amount       int64    `json:"amount" bson:"amount"`
	Days         int64    `json:"days" bson:"days"`
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
