package main

import (
	bootstrap "Zhigui/bootstap"
	"Zhigui/pkg/config"
	"Zhigui/router"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error

func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载.env文件，如 --env=testing 加载的是 .env.testing")
	flag.Parsed()
	config.InitConfig(env)
	// new 一个gin engine 实例
	r := gin.New()
	// 初始化路由绑定
	router.SetRoute(r)

	// 运行服务
	err := r.Run(":8848")
	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
	// 初始化	DB
	bootstrap.SetupDB()

	config.InitConfig(env)

	// 初始化Logger
	bootstrap.SetupLogger()

	// 设置gin的运行模式，支持 debug,release,test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	gin.SetMode(gin.ReleaseMode)

	// 初始化 Redis
	bootstrap.SetupRedis()
	
}
