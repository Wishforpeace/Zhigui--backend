package model

import "github.com/jinzhu/gorm"

 type User struct {
	 gorm.Model
	 StudentID string	`json:"student_id"`
	 PassWord string `json:"password" `
	 Name string
	 Info Info
 }


type Info struct {
	gorm.Model
	TasksPublished  []Task
	TasksFinished []Task
}