package controller_admin_storeaccess

import (
	"easyCardGen/model"
	serviceadminaccess "easyCardGen/service/admin/access"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddAccess 用于 添加商店权限
func AddAccess(c *gin.Context) {
	var storeAccess model.StoreAccess

	err := c.ShouldBindJSON(&storeAccess)
	if err != nil {
		return
	}

	//执行服务层
	err = serviceadminaccess.AddAccess(storeAccess)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
