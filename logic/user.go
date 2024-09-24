package logic

import (
	"web_framework/dao/mysql"
	"web_framework/pkg/snowflake"

)

func Register() {
	// 查询用户是否存在
	mysql.QueryUser()
	//1.生成Uid
	snowflake.GenID()
	//2.插入数据库
	mysql.InsertUser()
}