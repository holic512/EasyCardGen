package serviceadmin

import (
	"easyCardGen/model"
	repositoryusers "easyCardGen/repository/users"

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

	err := repositoryusers.FindUser(user)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.GetHashedPassword(user.Password)
	if err != nil {
		return errors.New("密码转换为哈希错误")
	}

	user.Password = hashedPassword

	err = repositoryusers.CreateUser(user)
	if err != nil {
		return errors.New("user表插入数据库错误")
	}

	return nil
}

func GetUserCount() (int64, error) {
	count, err := repositoryusers.GetUserCount()
	return count, err
}

func GetUsersInfo(currentPage int64, pageSize int64) ([]model.User, error) {
	usersInfo, err := repositoryusers.GetUserInfo(currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	return usersInfo, nil
}
