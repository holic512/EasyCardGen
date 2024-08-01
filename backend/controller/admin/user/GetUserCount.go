package controller_admin_user

import (
	service_admin_user "easyCardGen/service/admin/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserCount 用于 查询 用户数量 控制器
func GetUserCount(c *gin.Context) {
	count, err := service_admin_user.GetCount()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
		return
	}
	c.JSON(200, gin.H{"count": count})
}
