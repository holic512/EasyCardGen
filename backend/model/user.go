package model

import "time"

// User 用于存储用户类
type User struct {
	// 用户唯一ID（主键）
	ID uint `gorm:"primaryKey"`

	/*
		Username
		用户名（唯一约束，可能需要额外的索引或唯一约束设置）
		长度限制 4 - 12
	*/
	Username string `gorm:"type:varchar(12);uniqueIndex"`

	// 密码 (哈希加密)
	Password string `gorm:"type:varchar(255)"`

	// Email 用户邮箱
	Email string `gorm:"type:varchar(255)"`

	// IfEmail 邮箱验证
	IfEmail bool

	// Phone 用户电话
	Phone string `gorm:"type:varchar(255)"`

	// IfPhone 是否验证电话
	IfPhone bool

	// UserType 用户身份 (用户 0 商户 1 管理员 2)
	UserType string `gorm:"type:char(1)"`

	// State 状态 (禁用 0  启用 1 封禁 2)
	State string `gorm:"type:char(1)"`

	// 账号创建时间
	CreatedAt time.Time
	// 账号上次更新的时间（通常不是“上次登录的时间”，但如果你有特殊逻辑来更新它也可以）
	UpdatedAt time.Time
}
