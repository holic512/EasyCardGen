package service_admin_access

import (
	"easyCardGen/model"
	repositoryaccess "easyCardGen/repository/store_access"
)

// GetList 用于获取 商店权限列表
func GetList() ([]model.AccessListDto, error) {

	// 调用数据查询
	accessList, err := repositoryaccess.GetAccessList()
	if err != nil {
		return nil, err
	}
	return accessList, nil

}
