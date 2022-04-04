package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	_ "github.com/thedevsaddam/govalidator"
)

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// @Summary 登录
// @Tags user
// @Description 邮箱注册登录
// @Accept application/json
// @Product application/json
// @Param object body model.User true "登录用户信息"
// @Success 200 {object} Token "将studetn_id作为token保留"
// @Success 200 {object} handler.Response "{"msg":将studen_id作为token保留}"
// @Failure 401 {object} errno.Errno "{"error_code":"10001", "message":"Password or account wrong."} 身份认证失败 重新登录"
// @Failure 400 {object} errno.Errno "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /login [post]
func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
