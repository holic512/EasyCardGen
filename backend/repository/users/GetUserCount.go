package repositoryusers

import (
	"easyCardGen/config"
	"easyCardGen/model"
)

// GetUserCount 用于 获取 用户数量
func GetUserCount() (int64, error) {
	db := config.GetDB()
	var count int64
	result := db.Model(model.User{}).Count(&count)
	return count, result.Error
}
