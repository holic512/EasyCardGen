package admin

import (
	"easyCardGen/model"
	"easyCardGen/repository/users"
	"easyCardGen/utils"
	"errors"
)

func AddUser(user model.User) error {

	//验证数据规范性
	if user.Username == "" || user.Password == "" {
		return errors.New("USERNAME_PASSWORD_EMPTY")
	}
	if len(user.Username) < 4 || len(user.Username) > 12 {
		return errors.New("username should be between 4 and 12 characters in length")
	}
	if len(user.Password) < 8 || len(user.Password) > 16 {

		return errors.New("password should be between 8 and 16 characters in length")
	}

	// 检测 用户名密码重复
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
