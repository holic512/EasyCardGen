package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

import "easyCardGen/model"

var db *gorm.DB

func Init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"m13245jia",
		"127.0.0.1",
		"3306",
		"easycardgen",
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	} // 数据库连接成功时的处理逻辑

	fmt.Println("Success to connect database")

	// 自动迁移
	migrate()
}

// GetDB 用来返回已连接的数据库对象
func GetDB() *gorm.DB {
	return db
}

func migrate() {
	// 这里进行自动迁移

	//迁移 用户 数据表
	err := db.AutoMigrate(model.User{})
	if err != nil {
		return
	}

	//迁移 商店权限 数据表
	err = db.AutoMigrate(model.StoreAccess{})
	if err != nil {
		return
	}

	fmt.Println("Success to migrate")
}
