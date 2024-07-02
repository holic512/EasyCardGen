package repositoryusers

import (
	"easyCardGen/config"
	"easyCardGen/model"
	"errors"
)

// FindUser 用于查询是否有重复的username email phone
func FindUser(user model.User) error {
	db := config.GetDB()

	var count int64
	db.Model(model.User{}).Select("id").Where("username = ?", user.Username).Count(&count)
	if count != 0 {
		return errors.New("USERNAME_DUPLICATE_ERROR")
	}

	db.Model(model.User{}).Select("id").Where("email = ?", user.Email).Count(&count)
	if count != 0 {
		return errors.New("EMAIL_DUPLICATE_ERROR")
	}

	db.Model(model.User{}).Select("id").Where("phone = ?", user.Phone).Count(&count)
	if count != 0 {
		return errors.New("PHONE_DUPLICATE_ERROR")
	}

	return nil
}
