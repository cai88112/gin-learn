package router

import (
	"ginLearn/controller"
	"ginLearn/middleware"
	"github.com/gin-gonic/gin"
)


func UserRout(r *gin.Engine)  {
	urout := r.Group("/user").Use(middleware.TokenCheck())
	{
		urout.GET("/getInfo/:id",controller.GetInfo)
		urout.GET("/getList",controller.GetList)
		urout.PUT("/updateUser",controller.UpdateUser)
	}
}
