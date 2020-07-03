package server

import (
	"fmt"
	"ginLearn/dbdao"
	"ginLearn/global"
	"github.com/spf13/viper"
)

func init()  {
	fmt.Println("server init config,db,redis....")
	initConfig()
	initDb()
	initRedis()
}

func initConfig(){
	v := viper.New()
	v.AddConfigPath("./config/")
	v.SetConfigName("config")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	global.Vp = v
}
func initDb()  {
	dbdao.InitDb()
}
func initRedis()  {
	dbdao.Redis()
}
