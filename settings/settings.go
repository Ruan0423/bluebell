package settings

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

//conf 全局变量，保存conffig的配置信息
var Conf = new(APPConfig)
//使用结构体保存配置信息
type APPConfig struct {

	*APP `mapstructure:"app"`
	*Logconfig `mapstructure:"log"`
	*Mysqlconfig `mapstructure:"mysql"`
	*Redisconfig `mapstructure:"redis"`
	
}
type APP struct {
	Name string `mapstructure:"name"`
	Model string `mapstructure:"model"`
	Version string `mapstructure:"version"`
	Port int `mapstructure:"port`
}

type Logconfig struct {
    Filename string `mapstructure:"filename"`
    Level string `mapstructure:"level"`
	MaxSize int `mapstructure:"maxsize"`
	MaxBackups int `mapstructure:"maxbackups"`
	MaxAge int `mapstructure:"maxage"`
}

type Mysqlconfig struct {

	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Table string `mapstructure:"table"`
	Field string `mapstructure:"field"`
	TimeField string `mapstructure:"time_field"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type Redisconfig struct {

	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db string `mapstructure:"db"`
}
    // 全局变量

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

	if err = viper.Unmarshal(Conf); err != nil {

		fmt.Println("config init erro:",err)
		return err
	}
	fmt.Println("test config",Conf)

	//监听配置文件
	viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("Config file changed:", e.Name)
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Println("config init erro:",err)
		}
    })
	return nil

}