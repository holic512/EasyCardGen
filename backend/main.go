package main

import (
	"easyCardGen/router"
	"github.com/gin-gonic/gin"
)
import "easyCardGen/config"

func main() {

	//	程序入口

	server := gin.Default()
	router.SetupRoutes(server)
	err := server.Run(config.ServerAddress)
	if err != nil {
		return
	}

}
