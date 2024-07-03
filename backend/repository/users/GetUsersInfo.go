package repositoryusers

import (
	"easyCardGen/config"
	"easyCardGen/model"
	"fmt"
)

// GetUsersInfo 用于获取 对应页数 用户信息
func GetUsersInfo(currentPage int64, pageSize int64) ([]model.User, error) {
	var users []model.User
	offset := (currentPage - 1) * pageSize
	limit := pageSize

	db := config.GetDB()

	// 查询总记录数
	var totalRows int64
	if err := db.Model(&model.User{}).Count(&totalRows).Error; err != nil {

		return nil, err
	}

	// 调整页码，确保不超过实际页数
	if currentPage*pageSize >= totalRows {
		limit = totalRows - ((currentPage - 1) * pageSize)
	}

	if err := db.Model(&model.User{}).
		Select("ID", "Username", "Phone", "Email", "UserType", "State").
		Offset(int(offset)).Limit(int(limit)).Find(&users).Error; err != nil {
		fmt.Println("查询记录错误", err)
		return nil, err
	}

	return users, nil
}
