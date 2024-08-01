package service_admin_user

import (
	"easyCardGen/model"
	repositoryusers "easyCardGen/repository/users"
)

// GetList 用于获取用户列表
func GetList() ([]model.UserListDto, error) {
	//	调用 数据库查询
	userList, err := repositoryusers.GetUserList()
	if err != nil {
		return nil, err
	}
	return userList, nil
}
