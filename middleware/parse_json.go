package middleware

import (
	"github.com/gin-gonic/gin"
)

func ParseJson()  gin.HandlerFunc{
	 return func(c *gin.Context) {
	 	json := c.PostForm("inputJson")
	 	if json != ""{
	 		c.Set("inputJson",json)
		}
		 c.Next()
	 }
}
