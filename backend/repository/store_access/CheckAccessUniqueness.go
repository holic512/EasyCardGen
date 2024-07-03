package repositoryaccess

import (
	"easyCardGen/config"
	"easyCardGen/model"
	"errors"
)

// CheckAccessUniqueness 用于查询 商店权限 一致性
func CheckAccessUniqueness(storeAccess model.StoreAccess) error {
	db := config.GetDB()

	var count int64

	// 检测 权限 名称是否重复
	db.Model(&model.StoreAccess{}).Where("access_name = ?", storeAccess.AccessName).Count(&count)
	if count != 0 {
		return errors.New("ACCESS_NAME_DUPLICATE_ERROR")
	}

	return nil
}
