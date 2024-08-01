package service_admin_user

import repositoryusers "easyCardGen/repository/users"

// GetCount 用于获取用户数目
func GetCount() (int64, error) {
	return repositoryusers.GetUserCount()
}
