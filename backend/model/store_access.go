package model

import "time"

// StoreAccess 用于 商店权限
type StoreAccess struct {
	ID uint `gorm:"primaryKey"`

	AccessName string `gorm:"type:varchar(255)"`

	Type string `gorm:"type:varchar(255)"`

	MaxClassCount uint

	MaxProductCount uint

	Weight uint

	WithdrawalFee float64

	DiscountCode bool

	MembershipCard bool

	KeywordReply bool

	PlatformService bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
id
权限名称
颜色样式
最大分类数
最大商品数
推荐位权重
提现手续费
折扣码功能
会员卡功能
关键词回复
平台客服
创建时间
最近更新时间
*/

// AccessListDto 用于处理 获取用户列表 的 数据传输
type AccessListDto struct {
	ID         uint
	AccessName string
	Type       string
}
