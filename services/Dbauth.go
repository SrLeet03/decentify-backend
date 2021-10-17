package services

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Con *gorm.DB

type Cred struct {
	AuthSource string
	Username   string
	Password   string
}

func ConnectionwithDatabase() {

	con, err := gorm.Open("mysql", "admin"+":"+"happilyeveradmin"+"@"+"tcp(127.0.0.1:3306)"+"/"+"decentify")
	Con = con
	if err != nil {
		fmt.Println("Error while connectiong with sql databases", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error while connectiong with nosql databases", err.Error())
		return
	}

	fmt.Println("connection with sql and mongo the databases establibshed")
	Migrate(Con)
}
