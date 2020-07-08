package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var(
	Db *gorm.DB
	Vp *viper.Viper
	Redis *redis.Client
)
