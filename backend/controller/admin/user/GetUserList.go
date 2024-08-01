package controller_admin_user

import (
	serviceadmin "easyCardGen/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserList 用于查询用户列表(用于添加 其他信息 的索引外键)
func GetUserList(c *gin.Context) {
	//调用 查询 服务层
	userList, err := serviceadmin.GetUserList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
	}

	c.JSON(200, gin.H{"userList": userList})

}
