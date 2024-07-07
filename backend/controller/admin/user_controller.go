package controlleradmin

import (
	serviceadmin "easyCardGen/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
import "easyCardGen/model"

// AddUser 用于 添加 用户 控制器
func AddUser(c *gin.Context) {

	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := serviceadmin.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User added successfully"})
}

// GetUserCount 用于 查询 用户数量 控制器
func GetUserCount(c *gin.Context) {
	count, err := serviceadmin.GetUserCount()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
		return
	}
	c.JSON(200, gin.H{"count": count})
}

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

// GetUserList 用于查询用户列表(用于添加 其他信息 外键)
func GetUserList(c *gin.Context) {
	//调用 查询 服务层
	userList, err := serviceadmin.GetUserList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
	}

	c.JSON(200, gin.H{"userList": userList})

}
