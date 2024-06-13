package model

import "time"

//用于存储用户类

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Password  string
	Email     string
	Phone     string
	State     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
