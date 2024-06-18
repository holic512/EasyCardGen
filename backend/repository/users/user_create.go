package users

import (
	"easyCardGen/config"
	"easyCardGen/model"
)

func UserCreate(user model.User) error {
	db := config.GetDB()
	db.Create(&user)
	return nil
}
