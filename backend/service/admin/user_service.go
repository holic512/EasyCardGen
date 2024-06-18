package admin

import (
	"easyCardGen/model"
	"easyCardGen/repository/users"
	"easyCardGen/utils"
	"errors"
)

func AddUser(user model.User) error {
	//todo 验证部分数据 是否唯一

	hashedPassword, err := utils.GetHashedPassword(user.Password)
	if err != nil {
		return errors.New("密码转换为哈希错误")
	}
	user.Password = hashedPassword

	err = users.UserCreate(user)
	if err != nil {
		return errors.New("user表插入数据库错误")
	}
	return nil
}
