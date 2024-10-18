package controlle

import (
	"web_framework/logic"
	"web_framework/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//1.获取表单参数表
	param := new(models.ParamLogin)
	if err:= c.ShouldBindJSON(param); err !=nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2 业务处理（登录验证,生成token）
	if token,err:=logic.Login(param);err != nil {
		ResponseErrorwithMsg(c, CodeServerBusy, err.Error())
		return
	} else {

		// 3 响应结果
		ResponseSuccess(c, token)

	}
	
}