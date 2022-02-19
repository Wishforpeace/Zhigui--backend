package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Content string `gorm:"size:300"`
	Award string 
	State string
	
}