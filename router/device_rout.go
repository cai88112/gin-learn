package router

import (
	"ginLearn/controller"
	"github.com/gin-gonic/gin"
)

func DeviceRout(r *gin.Engine) {
	device :=r.Group("/init")
	{
			device.POST("/login",controller.DeviceLogin)
	}
}
