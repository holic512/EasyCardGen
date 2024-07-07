package model

import "time"

// User 用于存储用户类
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
	IfEmail  bool
	Phone    string `gorm:"type:varchar(255)"`
	IfPhone  bool
	UserType string `gorm:"type:varchar(255)"`
	State    string `gorm:"type:varchar(255)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

/*

用户唯一ID（主键）
用户名（唯一约束）
密码 (哈希加密)
邮箱
是否验证邮箱
电话
是否验证电话
用户身份 用户 商户
账号创建时间
账号上次登录的时间
*/

// UserListDto 用于处理 获取用户列表 的 数据传输
type UserListDto struct {
	ID       uint
	Username string
}
