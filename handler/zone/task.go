package zone

import (
	"Zhigui/model/task"
	"github.com/gin-gonic/gin"
)

// @Summary	 查看分区
// @Tags zone
// @Description 显示所需的分区
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param	
// @Success 200 {object} []model.Collection "{"msg":"获取成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Router /user/colletion [get]
func GetZones(c *gin.Context) {
	task := task.Task{}
	task := task.
}
