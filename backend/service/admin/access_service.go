package serviceadmin

import (
	"easyCardGen/model"
	repositoryaccess "easyCardGen/repository/store_access"
	"errors"
)

// AddAccess 用于添加商店权限
func AddAccess(storeAccess model.StoreAccess) error {

	// 检查数据规范性
	if storeAccess.AccessName == "" || storeAccess.Type == "" {
		return errors.New("ACCESS_NAME_TYPE_EMPTY")
	}

	// 检查数据是否重复
	err := repositoryaccess.CheckAccessUniqueness(storeAccess)
	if err != nil {
		return err
	}

	// 数据插入数据库
	err = repositoryaccess.CreateAccess(storeAccess)
	if err != nil {
		return err
	}
	// success
	return nil
}

// GetAccessList 用于获取 商店权限列表
func GetAccessList() ([]model.AccessListDto, error) {

	// 调用数据查询
	accessList, err := repositoryaccess.GetAccessList()
	if err != nil {
		return nil, err
	}
	return accessList, nil

}
