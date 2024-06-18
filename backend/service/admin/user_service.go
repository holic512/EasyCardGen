package admin

import (
	"easyCardGen/model"
	"easyCardGen/repository/users"
	"easyCardGen/utils"
	"errors"
)

func AddUser(user model.User) error {

	err := users.FindUser(user)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.GetHashedPassword(user.Password)
	if err != nil {
		return errors.New("密码转换为哈希错误")
	}

	user.Password = hashedPassword

	err = users.CreateUser(user)
	if err != nil {
		return errors.New("user表插入数据库错误")
	}

	return nil
}
