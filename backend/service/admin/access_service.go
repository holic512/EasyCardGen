package serviceadmin

import (
	"easyCardGen/model"
	repositoryaccess "easyCardGen/repository/store_access"
	"errors"
)

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
