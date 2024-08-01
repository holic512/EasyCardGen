package controller_admin_user

import (
	"easyCardGen/model"
	serviceadmin "easyCardGen/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUsersInfo 用于 查询 用户信息(显示在表格中) 控制器
func GetUsersInfo(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	//pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	var usersInfo []model.User
	var err error
	usersInfo, err = serviceadmin.GetUsersInfo(int64(currentPage), int64(10))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
	}

	c.JSON(200, gin.H{"userData": usersInfo})
}
