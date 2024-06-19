package admin

import (
	"easyCardGen/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "easyCardGen/model"

// AddUser 用于用户逻辑处理
func AddUser(c *gin.Context) {

	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := admin.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User added successfully"})
}
