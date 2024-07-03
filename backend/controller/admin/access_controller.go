package controlleradmin

import (
	"easyCardGen/model"
	serviceadmin "easyCardGen/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAccess(c *gin.Context) {
	var storeAccess model.StoreAccess

	err := c.ShouldBindJSON(&storeAccess)
	if err != nil {
		return
	}

	//执行服务层
	err = serviceadmin.AddAccess(storeAccess)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
