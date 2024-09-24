package controlle

import (
	"fmt"
	// "web_framework/logic"
	modlels "web_framework/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	//1.参数接受和校验
	param := new(modlels.ParamRegister)

	if err :=c.ShouldBindJSON(param);err != nil {
		// zap.L().Error("register handler failed", zap.Error(err))
		c.JSON(200,gin.H{
			"msg":"请求参数有误！",
		})
		return
	}

	fmt.Println(param)

	if param.Password != param.Repassword {
		zap.L().Error("两次输入的密码不一样！")
		c.JSON(200,gin.H{
			"msg":"两次输入的密码不一样！",
		})
		return
	}
	if len(param.Password) == 0 || len(param.Username) ==0{
		zap.L().Error("参数不能为空！")
        c.JSON(200,gin.H{
            "msg":"参数不能为空！",
        })
        return
	}


	//2.业务处理（数据库操作，插入数据）
	// logic.Register()

	//3. 响应结果
	c.JSON(200, "registe success!")
}