package controlle

import (
	"web_framework/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 社区相关的路由处理

func CommunityHandler(c *gin.Context) {
	// 获取社区分类
	data , err :=logic.GetCommunitiList()
	if err!=nil {
		zap.L().Error("logic.GetComminityList err:",zap.Error(err))
		ResponseError(c,CodeServerBusy)
	}
	ResponseSuccess(c,data)
}