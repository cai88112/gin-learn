package controller

import (
	"encoding/json"
	"fmt"
	"ginLearn/common"
	"ginLearn/global"
	"ginLearn/model"
	"ginLearn/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func Login(c *gin.Context)  {
	//username := c.PostForm("userName")
	//pwd := c.PostForm("pwd")
	id := c.PostForm("id")
	var device model.Device
	 if global.Db.Table("member").Where("id = ?",id).First(&device).RecordNotFound(){
	 	response.FailWithMessage("找不到用户",c);
		 return
	 }else{
		 rtjson,_:=json.Marshal(device)
		 fmt.Println("byte format:%s",rtjson)
		 st, _ := strconv.Unquote(string(rtjson))
		 fmt.Println("string format:%s",st)
		 response.OkDetailed(map[string]interface{}{
			 "data":device,
		 },"查询成功",c)
	 }

	if err:=json.Unmarshal([]byte(c.PostForm("inputJson")),&device);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg": "error:"+err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "解析成功",
		"parseJson":device,
	})
}
func GetInfo(c *gin.Context){
		tokenString := c.PostForm("token")
		claims := &common.Claim{}
		flag,claims := common.CheckToken(tokenString)
		if !flag{
			c.JSON(400, gin.H{
				"code": 401,
				"msg": "token valid error",
			})
			return
		}
	c.JSON(200, gin.H{
		"code": 200,
		"device": claims,
	})
}
