package router

import (
	"Zhigui/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断email是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
		}
	}
}
