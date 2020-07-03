package router

import (
	"github.com/gin-gonic/gin"
)


func GetAllRoutes(r *gin.Engine) *gin.Engine{
		DeviceRout(r)
		UserRout(r)
		return r
}