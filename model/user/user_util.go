package user

import (
	"Zhigui/model"
	"Zhigui/pkg/database"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
