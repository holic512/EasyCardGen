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

	//验证数据规范性
	if newUser.Username == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password fields are empty"})
		return
	}
	if len(newUser.Username) < 4 || len(newUser.Username) > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username should be between 4 and 12 characters in length"})
		return
	}
	if len(newUser.Password) < 8 || len(newUser.Password) > 16 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password should be between 8 and 16 characters in length"})
		return
	}

	err := admin.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User added successfully"})
}
