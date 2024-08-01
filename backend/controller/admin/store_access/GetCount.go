package controller_admin_storeaccess

import (
	repositoryaccess "easyCardGen/repository/store_access"
	"github.com/gin-gonic/gin"
)

// GetCount 用于获取 商店权限的 数量
func GetCount(c *gin.Context) {
	count, err := repositoryaccess.GetCount()
	if err != nil {
		c.JSON(404, gin.H{"error_code": err.Error()})
	}
	c.JSON(200, gin.H{"accessCount": count})
}
