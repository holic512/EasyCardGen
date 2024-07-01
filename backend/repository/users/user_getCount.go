package repository_users

import (
	"easyCardGen/config"
	"easyCardGen/model"
)

// GetUserCount 用于 创建 用户信息
func GetUserCount() (int64, error) {
	db := config.GetDB()
	var count int64
	result := db.Model(model.User{}).Count(&count)
	return count, result.Error
}
