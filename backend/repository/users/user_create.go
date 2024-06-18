package users

import (
	"easyCardGen/config"
	"easyCardGen/model"
)

// CreateUser 用于 创建 用户信息
func CreateUser(user model.User) error {
	db := config.GetDB()
	db.Create(&user)
	return nil
}
