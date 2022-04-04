package user

import (
	"Zhigui/model/task"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	EMail           string `json:"e_mail"`
	PassWord        string `json:"password" `
	Name            string `json:"name" gorm:"type:varchar(255);not null;unique"`
	ActivationToken string `gorm:"varchar(255)"`
	Info            string
	TaskPublished   []task.Task
	TaskFinished    []task.Task
}
