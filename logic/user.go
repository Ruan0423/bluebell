package logic

import (
	"errors"
	"fmt"
	"web_framework/dao/mysql"
	"web_framework/models"
	"web_framework/pkg/jwt"
	"web_framework/pkg/snowflake"

	"go.uber.org/zap"
)

// register注册用户业务处理
func Register(param *models.ParamRegister) error{
	// 查询用户是否存在
	if mysql.Check(param.Username) {
		return errors.New("用户已存在！")
	}
	//1.生成Uid
	Uid :=snowflake.GenID()
	//生成用户
	fmt.Println("插入数据库出错")
	user := &models.User{
		USER_ID: Uid,
		Username: param.Username,
		Password: param.Password,
	}

	//2.插入数据库
	err :=mysql.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func Login(param *models.ParamLogin) (*models.User, error) {

	//1.验证用户是否存在
	if !mysql.Check(param.Username) {
		zap.L().Error("用户不存在，",zap.Any("username",param.Username))
		return nil , errors.New("用户或密码错误！")
	}

	//2.验证密码
	user := &models.User{
		Username: param.Username,
		Password: param.Password,
	}
	if !mysql.Login(user) {
		zap.L().Error("密码不正确，",zap.Any("username",param.Username),zap.Any("password",param.Password))
		return nil,errors.New("用户或密码错误！")
	} else {
		token,err := jwt.GenToken(user.Username, user.USER_ID)
		user.Token = token
		return user,err
	}

}