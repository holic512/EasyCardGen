package model

import "time"

//用于存储用户类

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
	Phone    string
	userType string
	State    string

	CreatedAt time.Time
	UpdatedAt time.Time
}

/*

用户唯一ID（主键）
用户名（唯一约束）
密码 (哈希加密)
邮箱
电话
账号状态 正常 停用 封禁
用户身份 用户 商户
账号创建时间
账号上次登录的时间
*/
