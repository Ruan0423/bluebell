package mysql

import (
	"fmt"
	"web_framework/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)
var Db *gorm.DB

func Init(cfg *settings.Mysqlconfig) (err error) {


	dsn :=fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	cfg.Username,
	cfg.Password,
	cfg.Host,
	cfg.Port,
	cfg.Database,
)
	fmt.Println(dsn)
	Db,err = gorm.Open("mysql",dsn)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	

	return nil
} 