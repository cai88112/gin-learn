package controller

import (
	"ginLearn/global"
	"ginLearn/model"
	"ginLearn/response"
	"ginLearn/util"
	"github.com/gin-gonic/gin"
	"time"
)


func GetInfo(c *gin.Context)  {
	//username := c.PostForm("userName")
	//pwd := c.PostForm("pwd")
	id := c.Param("id")
	var user model.User
	 if err :=global.Db.Table("member").Where("id = ?",id).First(&user).Error;err != nil{
	 	response.FailWithMessage("找不到用户",c)
		 return
	 }else{
		 response.Result("0000",user,"操作成功",c)
	 }
}
func GetList(c *gin.Context){
		var users []model.User
		if err:=global.Db.Table("member").Order("id desc").Find(&users).Error;err!=nil{
			response.FailWithMessage("查询失败",c)
			return
		}
		response.Result("0000",users,"操作成功",c)
}

func  UpdateUser(c *gin.Context){
	var user model.User
	if err:=c.ShouldBind(&user);err!=nil{
		response.FailWithMessage("参数错误",c)
		return
	}
	 updateD := util.Time(time.Now())

	if err:=global.Db.Table("member").Model(&user).Updates(model.User{NickName: user.NickName, Mobile: user.Mobile, UpdateD: updateD}).First(&user).Error;err!=nil{
		response.FailWithMessage("修改失败",c)
		return
	}
	response.Result("0000",user,"操作成功",c)
}
