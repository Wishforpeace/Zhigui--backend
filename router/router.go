package router

import (
	"Zhigui/handler/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router()*gin.Engine{
	r := gin.New()
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"The incorrect API router")
	})
	//登录
	auth := r.Group("/api/v1/login")
	{
		auth.POST("",handler.Login)
	}
	//
	return r
}	