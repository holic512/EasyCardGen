package service_admin_user

import (
	"easyCardGen/model"
	repositoryusers "easyCardGen/repository/users"
)

// GetData 用于 获取用户详细数据
func GetData(currentPage int64, pageSize int64) ([]model.User, error) {
	usersInfo, err := repositoryusers.GetUsersInfo(currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	return usersInfo, nil
}
