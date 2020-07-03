package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var(
	Db *gorm.DB
	Vp *viper.Viper
	Redis *redis.Client
)
