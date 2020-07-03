package router

import (
	"ginLearn/controller"
	"github.com/gin-gonic/gin"
)


func UserRout(r *gin.Engine)  {
	urout := r.Group("/user")
	{
		urout.POST("/login",controller.Login)
		urout.POST("/getInfo",controller.GetInfo)
	}
}
