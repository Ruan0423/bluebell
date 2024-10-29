package mysql

import (
	"database/sql"
	"errors"
	"web_framework/models"

	"go.uber.org/zap"
)

func Createpost(p *models.Post) error {
	sqlstr := "insert into post (post_id,author_id,community_id,title,content) values (?,?,?,?,?)"

	if _,err:= db.Exec(sqlstr,p.ID,p.AuthorID,p.CommunityID,p.Title,p.Content); err!=nil {
		zap.L().Error("Error inserting post",zap.Error(err))
		return err
	}
	return nil
}

func GetPOstByID(postid int64) (postdata *models.Post, err error) {
	sqlstr := "select post_id,author_id,community_id,title,content,create_time,update_time from post where post_id =?"
	postdata = new(models.Post)
	err = db.Get(postdata,sqlstr,postid)
	
	if err!=nil {
		if err == sql.ErrNoRows{
			zap.L().Error("没有此文章",zap.Any("文章id:",postid))
			err = nil
			postdata = nil
			return 
		}
		zap.L().Info("文章id查询帖子数据出错：",zap.Error(err))
	}
	return
}

func GetPostList(offset , pagesize int64) ([]*models.Post , error){
	sqlstr := "select post_id,author_id,community_id,title,content,create_time,update_time from post limit ?,?"

	postlist := make([]*models.Post, 0,pagesize)
	err := db.Select(&postlist, sqlstr, offset,pagesize)
	if err!=nil {
		zap.L().Error("获取帖子列表失败GetPostList：", zap.Error(err))
		err = errors.New("获取帖子列表失败！")
	}
	return postlist, err
}