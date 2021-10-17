package services

import (
	"github.com/SrLeet03/decentify-backend/models"
	"github.com/jinzhu/gorm"
)

func Migrate(Conn *gorm.DB) {

	Conn.AutoMigrate(models.RegisReq{})
}
