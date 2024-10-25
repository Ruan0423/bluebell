package controlle

import (
	"fmt"
	"web_framework/logic"
	"web_framework/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	// "go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	//1.参数接受和校验
	param := new(models.ParamRegister)

	if err :=c.ShouldBindJSON(param);err != nil {
		zap.L().Error("register handler failed", zap.Error(err))
		ResponseErrorwithMsg(c, CodeInvalidParam, err.Error())
		return
	}



	//2.业务处理（数据库操作，插入数据）
	if err:=logic.Register(param);err!=nil {
		ResponseErrorwithMsg(c, CodeServerBusy, err.Error())
        return     
	}

	//3. 响应结果
	ResponseSuccess(c, "注册成功")
} 

func Login(c *gin.Context) {
	//1.获取表单参数表
	param := new(models.ParamLogin)
	if err:= c.ShouldBindJSON(param); err !=nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2 业务处理（登录验证,生成token）
	if user,err:=logic.Login(param);err != nil {

		ResponseErrorwithMsg(c, CodeServerBusy, err.Error())
		return
	} else {

		// 3 响应结果
		ResponseSuccess(c, gin.H{
			"user_id": fmt.Sprintf("%d",user.USER_ID),
			"user_name": user.Username,
			"token": user.Token,
		})

	}
	
}