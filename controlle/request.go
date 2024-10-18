package controlle

import (
	"errors"

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