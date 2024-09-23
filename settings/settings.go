package settings

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

func Init() (err error){
	// 使用viper加载配置。
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err!= nil {
        fmt.Println("config init erro:",err)
		return err
    }

	//监听配置文件
	viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("Config file changed:", e.Name)
    })
	return nil

}