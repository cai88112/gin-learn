package dbdao

import (
	"fmt"
	"ginLearn/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
	"time"
)


func InitDb() *gorm.DB {
	viper := global.Vp
	//drivername := viper.GetString("db.drivername")
	host := viper.GetString("db.host")
	username := viper.GetString("db.username")
	pwd := viper.GetString("db.pwd")
	database := viper.GetString("db.database")
	charset := viper.GetString("db.charset")
	port := viper.GetString("db.port")
	loc := viper.GetString("db.location")
	//"gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		pwd,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
		)
	//args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
	//	username,
	//	pwd,
	//	host,
	//	port,
	//	database,
	//	charset,
	//	url.QueryEscape(loc),
	//)
	var logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 120 * time.Millisecond,
		LogLevel:      logger.Info,
		Colorful:      true,
	})
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{Logger:logger})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
	fmt.Println("connect db success")
	//// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//db.DB().SetMaxIdleConns(10)
	//
	//// SetMaxOpenConns sets the maximum number of open connections to the database.
	//db.DB().SetMaxOpenConns(100)
	//
	//// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//db.DB().SetConnMaxLifetime(time.Minute)

	global.Db = db
	return db
}

