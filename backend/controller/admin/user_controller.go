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

	//todo 验证数据规范性

	err := admin.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{})
}
