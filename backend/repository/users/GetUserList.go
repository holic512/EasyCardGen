package repositoryusers

import (
	"easyCardGen/config"
	"easyCardGen/model"
)

// GetUserList 用来获取 用户 列表
func GetUserList() ([]model.UserListDto, error) {
	var userList []model.UserListDto

	db := config.GetDB()
	db.Model(model.User{}).
		Select("id,username").
		Scan(&userList)

	return userList, nil
}
