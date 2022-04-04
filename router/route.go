package router

import (
	"Zhigui/router/middlewares"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)
	// 注册API路由
	RegisterAPIRoutes(router)
	// 配置404路由
	setup404Handler(router)

}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理404请求
	router.NoRoute(func(c *gin.Context) {
		// 获取表头信息的Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是HTML
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			// 默认返回JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
func registerGloablMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

//r.NoRoute(func (c *gin.Context) {
//	c.String(http.StatusNotFound, "The incorrect API router")
//})
////登录
//auth := r.Group("/api/v1/auth")
//{
//{
//auth.POST("/login", handler.Login)
//
//}
//}
////学习区
//zone := r.Group("/api/v1/task").Use(middleware.Auth())
//{
//zone.GET("", handler.GetZones)
//zone.GET("/partition", handler.GetPartition)
//zone.GET("/partition/task", handler.GetTasks)
//zone.POST("/partition/task", handler.PublishTask)
//zone.POST("/partition/:task_id", handler.AcceptTask)
//}
////用户
//user := r.Group("api/v1/user").Use(middleware.Auth())
//{
//user.GET("", handler.GetInfo)
//
//}
////排行榜
//list := r.Group("/api/v1/list").Use(middleware.Auth())
//{
//list.GET("")
//}
//
//return r
