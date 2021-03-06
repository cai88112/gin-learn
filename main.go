package main

import (
	_ "ginLearn/server"
	"ginLearn/global"
	"ginLearn/router"
	"github.com/gin-gonic/gin"
)
func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	router.GetAllRoutes(r)
	r.Run(":"+global.Vp.GetString("server.port"))
}


