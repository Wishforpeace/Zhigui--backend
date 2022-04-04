package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Publisher string
	Receiver  string
	Publish   User   `gorm:"foreignkey:Name;references:Publisher"`
	Receive   User   `gorm:"foreignkey:Name;references:Receiver"`
	Content   string `gorm:"size:300"`
	Award     string `json:"award"`
	State     string
}
