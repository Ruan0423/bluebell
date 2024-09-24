package controlle

import (
	"fmt"
	// "web_framework/logic"
	modlels "web_framework/models"

	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	//1.参数接受和校验
	param := new(modlels.ParamRegister)

	if err :=c.ShouldBindJSON(param);err != nil {
		// zap.L().Error("register handler failed", zap.Error(err))
		c.JSON(200,gin.H{
			"msg":err.Error(),
		})
		return
	}

	fmt.Println(param)



	//2.业务处理（数据库操作，插入数据）
	// logic.Register()

	//3. 响应结果
	c.JSON(200, "registe success!")
} 