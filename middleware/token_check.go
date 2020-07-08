package middleware

import (
	"fmt"
	"ginLearn/common"
	"ginLearn/response"
	"github.com/gin-gonic/gin"
)

func TokenCheck()  gin.HandlerFunc{
	 return func(c *gin.Context) {
		 tokenString := c.Query("token")
		 if tokenString == ""{
			 response.FailWithMessage("token can't be null",c)
			 c.Abort()
			 return

		 }
		 claims := &common.Claim{}
		 flag,claims := common.CheckToken(tokenString)
		 if !flag{
		 	response.FailWithMessage("token valid error",c)
			 c.Abort()
			 return

		 }
		fmt.Println("middleware user claims:="+claims.Name)
		 c.Next()
	 }
}
