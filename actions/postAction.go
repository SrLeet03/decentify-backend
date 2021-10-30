package actions

import (
	"github.com/SrLeet03/decentify-backend/models"
	"github.com/jinzhu/gorm"
)

func CreatePost(post models.PostReg, Con *gorm.DB) (string, string) {
	var findid models.Posts
	Con.Last(&findid)
	//fmt.Println(findid)
	var savepost models.Posts

	savepost.PostId = findid.PostId
	savepost.Uuid = GenerateUUID()
	savepost.UserId = post.UserId
	savepost.Title = post.Title
	savepost.Tag = post.Tag
	savepost.CreationTime = post.CreationTime
	savepost.Amount = post.Amount
	savepost.Days = post.Days

	err := Con.Create(&savepost)

	if err.RowsAffected == 0 || err.Error != nil {
		return "", err.Error.Error()
	}

	var savebody models.PostBody
	savebody.Body = post.Body
	savebody.PostId = findid.PostId
	err = Con.Create(&savebody)

	if err.RowsAffected == 0 || err.Error != nil {
		return "", err.Error.Error()
	}
	return "success", ""
}
