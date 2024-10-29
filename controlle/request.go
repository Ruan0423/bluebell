package controlle

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErroInfo = errors.New("用户未登录")
const (
	UserIDkey = "user_id"
)

func GetUserId(c *gin.Context) (user_id int64, err error) {
	uid ,ok := c.Get(UserIDkey)
	if !ok {
		err = ErroInfo
		return
	}
	user_id ,ok = uid.(int64)
	if !ok {
		err = ErroInfo
		return
	}
	return
}

// GetPOstParam 获取帖子列表的参数pagenum , pageSize
func GetPOstParam(c *gin.Context)(int64, int64){
	pagenumstr := c.Query("page")
	pageSizestr := c.Query("size")

	var (
		pagenum  int64
		pageSize int64
		err      error
	)
	//转化参数为整数
	pagenum, err = strconv.ParseInt(pagenumstr, 10, 64)
	if err != nil {
		pagenum = 1
	}
	pageSize, err = strconv.ParseInt(pageSizestr, 10, 64)
	if err != nil {
		pageSize = 10
	}
	return pagenum, pageSize
}