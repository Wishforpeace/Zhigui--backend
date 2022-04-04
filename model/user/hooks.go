package user

import (
	"Zhigui/pkg/strings"
	"github.com/jinzhu/gorm"
)

func (u *User) BeforeCrete(tx *gorm.DB) (err error) {
	u.ActivationToken = strings.Rand(10)
	return
}
