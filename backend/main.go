package main

import (
	"easyCardGen/config"
	"easyCardGen/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	//初始化数据库
	config.Init()

	// 初始化 gin
	server := gin.Default()

	// 使用自定义配置的 CORS 中间件
	core := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	server.Use(cors.New(core))

	//初始化路由
	router.SetupRoutes(server)

	//启动服务
	err := server.Run(config.ServerAddress)
	if err != nil {
		return
	}

}
