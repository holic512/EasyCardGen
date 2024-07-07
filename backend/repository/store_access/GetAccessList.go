package repositoryaccess

import (
	"easyCardGen/config"
	"easyCardGen/model"
)

// GetAccessList 用于获取 创建新商店 所需权限数据
func GetAccessList() ([]model.AccessListDto, error) {
	db := config.GetDB()
	var accessList []model.AccessListDto
	db.Model(&model.StoreAccess{}).
		Select("id,access_name,type").
		Scan(&accessList)

	return accessList, nil
}
