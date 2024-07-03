package repositoryaccess

import (
	"easyCardGen/config"
	"easyCardGen/model"
)

// CreateAccess 用于创建 商店权限
func CreateAccess(storeAccess model.StoreAccess) error {
	db := config.GetDB()
	db.Create(&storeAccess)
	return nil
}
