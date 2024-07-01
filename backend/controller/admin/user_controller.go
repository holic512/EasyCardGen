package controller_admin

import (
	service_admin "easyCardGen/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
import "easyCardGen/model"

func AddUser(c *gin.Context) {

	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service_admin.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User added successfully"})
}

func GetUserCount(c *gin.Context) {
	count, err := service_admin.GetUserCount()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
		return
	}
	c.JSON(200, gin.H{"count": count})
}

func GetUsersInfo(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	//pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	var usersInfo []model.User
	var err error
	usersInfo, err = service_admin.GetUsersInfo(int64(currentPage), int64(10))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
	}

	c.JSON(200, gin.H{"userData": usersInfo})
}
