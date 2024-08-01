package controller_admin_storeaccess

import (
	serviceadmin "easyCardGen/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAccessList 用于 获取权限列表 (仅获取名称,id)
func GetAccessList(c *gin.Context) {

	//	调用服务层
	accessList, err := serviceadmin.GetAccessList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"accessList": accessList})

}
