package model

import (
	"Zhigui/model/user"
	"github.com/jinzhu/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&user.User{})
}
