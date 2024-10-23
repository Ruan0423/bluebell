package mysql

import (
	"database/sql"
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