package controller_admin_user

import (
	"easyCardGen/model"
	service_admin_user "easyCardGen/service/admin/user"

	"github.com/gin-gonic/gin"
	"net/http"
)

// AddUser 用于 添加 用户 控制器
func AddUser(c *gin.Context) {

	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service_admin_user.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User added successfully"})
}
