package repositoryaccess

import (
	"easyCardGen/config"
	"easyCardGen/model"
)

// GetCount 用于查询 StoreAccess 表中数目
func GetCount() (int, error) {
	db := config.GetDB()
	var count int64
	result := db.Model(model.StoreAccess{}).Count(&count)
	return int(count), result.Error
}
