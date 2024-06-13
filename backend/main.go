package main

import (
	"easyCardGen/db"
	"easyCardGen/model"
)

func main() {

	//初始化数据库
	db.Init()

	user := model.User{
		Username: "123",
		Password: "123",
		Email:    "123",
		Phone:    "123",
		State:    "1223",
	}

	sql := db.GetDB()
	sql.Create(&user)

	println(user.ID)

	//// 初始化 gin
	//server := gin.Default()
	//
	////初始化路由
	//router.SetupRoutes(server)
	//
	////启动服务
	//err := server.Run(config.ServerAddress)
	//if err != nil {
	//	return
	//}

}
