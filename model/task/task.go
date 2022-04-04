package task

import (
	"Zhigui/model/user"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Publisher string
	Receiver  string
	Publish   user.User `gorm:"foreignkey:Name;references:Publisher"`
	Receive   user.User `gorm:"foreignkey:Name;references:Receiver"`
	Content   string    `gorm:"size:300"`
	Award     string    `json:"award"`
	State     string
}
