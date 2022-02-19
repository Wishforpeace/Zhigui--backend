package main

import (
	"Zhigui/model"
	"Zhigui/router"
	"fmt"

	// "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error

func main() {
	dsn := "root:root&1234@tcp(127.0.0.1:3306)/Zhigui?charset=utf8mb4&parseTime=True&loc=Local"
	model.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	// link := "http://127.0.0.1:"+
	defer model.DB.Close()
	r := router.Router()
	err = r.Run(":6666")
	if err !=nil{
		panic(err)
	}
}
