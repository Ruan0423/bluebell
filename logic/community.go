package logic

import (
	"web_framework/dao/mysql"
	"web_framework/models"
)

func GetCommunitiList() (data []*models.Community, err error) {
	// 从数据库中获取社区列表
	data, err = mysql.GetComminityList()
	return data,err

}

func GetCommunityDetatl(id int64) (data *models.CommunityDetail, err error) {
	return mysql.GetComminityDetailByID(id)
}