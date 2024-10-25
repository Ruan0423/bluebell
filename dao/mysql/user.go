package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"web_framework/models"
	modlels "web_framework/models"

	"go.uber.org/zap"
)

//查询
func Check(username string) (bool){
	var count int

	sqlstr := "select count(*) from user where username =?"

	err :=db.Get(&count,sqlstr,username)
	fmt.Println(count)
	if err !=nil {
		zap.L().Error("用户检测失败",zap.Error(err))
		return true
	}
	if count >0 {
		return true
	}
	return false

}

//插入
func InsertUser(user *modlels.User) error{
	//对用户密码进行加密
	user.Password = Encrypt(user.Password)
	// fmt.Println(user.Password)
	sqlstr := "insert into user (user_id, username, password) values (?, ?,?)"
	_,err:=db.Exec(sqlstr,user.USER_ID,user.Username,user.Password)
	if err !=nil {
		zap.L().Error("插入数据失败",zap.Error(err))
		fmt.Println(err)
		return errors.New("用户注册失败！")
	}
	return nil
}

// 登录检查密码
func Login(user *models.User) bool {
	// user := new(models.User)
	oPassword := user.Password

	sqlx := "select user_id,username,password from user where username = ?"
	err := db.Get(user,sqlx,user.Username)
	if err!= nil {
		zap.L().Error("查询用户失败",zap.Error(err))
		fmt.Println(err)
		return false
	}

	if user.Password != Encrypt(oPassword) {
		return false
	}
	return true
}
//删除

func DeleteUser(){

}

//修改

func UpdateUser(){

}

func Encrypt(password string) string {

	hasher := md5.New()
	hasher.Write([]byte(password))
	hashInBytes := hasher.Sum(nil)

	// 将散列值转换为十六进制格式
	md5String := hex.EncodeToString(hashInBytes)
	return md5String
}