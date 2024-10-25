package controlle

import (
	"strconv"
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

func CommunityDetailHandler(c *gin.Context) {
	// gey param 
	idstr := c.Param("id")

	id , err := strconv.ParseInt(idstr, 10, 64)
	if err!=nil {
		zap.L().Error("参数转换出错",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}

	//获取社区详情

	data,err := logic.GetCommunityDetatl(id)
	if err!=nil {
		zap.L().Error("获取社区详细信息失败，：",zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,data)
}