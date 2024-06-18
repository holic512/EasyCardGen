package utils

import "golang.org/x/crypto/bcrypt"

// GetHashedPassword 用于得到密码对应的哈希值
func GetHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CompareHashAndPassword 用于验证密码哈希值是否匹配
func CompareHashAndPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
