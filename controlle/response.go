package controlle

import "github.com/gin-gonic/gin"

type ResponsData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, Code ResCode) {
	c.JSON(200 , &ResponsData{
		Code: Code,
		Msg: Code.Msg(),
		Data: nil,
	})

}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200 , &ResponsData{
		Code: CodeSuccess,
		Msg: CodeSuccess.Msg(),
		Data: data,
	})

}
func ResponseErrorwithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(200 , &ResponsData{
		Code: code,
		Msg: msg,
		Data: nil,
	})

}