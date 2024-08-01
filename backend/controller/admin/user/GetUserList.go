package controller_admin_user

import (
	service_admin_user "easyCardGen/service/admin/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserList 用于查询用户列表(用于添加 其他信息 的索引外键)
func GetUserList(c *gin.Context) {
	//调用 查询 服务层
	userList, err := service_admin_user.GetList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
	}

	c.JSON(200, gin.H{"userList": userList})

}
