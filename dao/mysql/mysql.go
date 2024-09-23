package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
)
var Db *gorm.DB

func Init() (err error) {


	dsn :=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	viper.GetString("mysql.username"),
	viper.GetString("mysql.password"),
	viper.GetString("mysql.host"),
	viper.GetString("mysql.port"),
	viper.GetString("mysql.database"),
)
	fmt.Println(dsn)
	Db,err = gorm.Open("mysql",dsn)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	

	return nil
} 