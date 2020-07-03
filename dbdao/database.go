package dbdao

import (
	"fmt"
	"ginLearn/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/url"
)


func InitDb() *gorm.DB {
	viper := global.Vp
	drivername := viper.GetString("db.drivername")
	host := viper.GetString("db.host")
	username := viper.GetString("db.username")
	pwd := viper.GetString("db.pwd")
	database := viper.GetString("db.database")
	charset := viper.GetString("db.charset")
	port := viper.GetString("db.port")
	loc := viper.GetString("location")
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		pwd,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db, err := gorm.Open(drivername, args)
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
	db.LogMode(true)
	global.Db = db
	return db
}

