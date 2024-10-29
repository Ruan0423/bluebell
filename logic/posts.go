package logic

import (
	"web_framework/dao/mysql"
	"web_framework/models"
	"web_framework/pkg/snowflake"
)

func CreatePost(p *models.Post) error {
	//生成帖子id
	post_id := snowflake.GenID()
	p.ID = post_id
	return mysql.Createpost(p)
}

//GetPostApi 通过帖子Id获取帖子详情
func GetPostApi(postid int64) (postdata *models.Post, err error) {
	return mysql.GetPOstByID(postid)
}

// GetAthorByPostid 通过作者id获取作者名字
func GetAthorByAthorid(athor_id int64) (athor string, err error) {
	return mysql.GetAthorByUserid(athor_id)
}

// 
func GetPostList(pagenum , pageSize int64) ([]*models.Post, error) {

	//计算偏移量
	offset := (pagenum-1) * pageSize

	//查询
	return mysql.GetPostList(offset, pageSize)
}