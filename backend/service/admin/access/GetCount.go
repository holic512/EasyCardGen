package service_admin_access

import repositoryaccess "easyCardGen/repository/store_access"

// GetCount 用于 获取权限 数目
func GetCount() (int, error) {
	return repositoryaccess.GetCount()
}
