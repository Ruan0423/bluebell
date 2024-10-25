package mysql

import (
	"database/sql"
	"errors"
	"web_framework/models"

	"go.uber.org/zap"
)

func GetComminityList() (data []*models.Community, err error) {

	sqlsrt := "select community_id , community_name from community"

	if err = db.Select(&data,sqlsrt) ; err != nil {
		if err == sql.ErrNoRows{
			zap.L().Warn("getcommunity data failed with no data", zap.Error(err))
			return nil, nil
		}
		zap.L().Error("get community data failed", zap.Error(err))
		return nil, err
	}
	return
}

func GetComminityDetailByID(id int64) (data *models.CommunityDetail, err error) {
	sqlstr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	data = new(models.CommunityDetail)

	if err=db.Get(data,sqlstr,id);err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("id错误") 
		}
		zap.L().Error("获取查询社区详情失败：",zap.Error(err))
	}
	return 
}