package controller

import (
	"ginLearn/common"
	"ginLearn/log"
	"ginLearn/response"
	"ginLearn/service"
	"github.com/gin-gonic/gin"
)
var mylog = log.GetLogger("device.log")

func DeviceLogin(c *gin.Context)  {
	username := c.PostForm("userName")
	pwd := c.PostForm("pwd")
	num := c.PostForm("deviceNum")
	err := service.ValidAccount(username,pwd)
	if err != nil{
		response.FailWithDetailed("1005",err.Error(),"账户或密码错误",c)
		return
	}

	device,err := service.ValidDevice(num)
	if err != nil{
		response.FailWithDetailed("1015","","当前设备不合法",c)
		return
	}else {
		token, err := common.GetToken(device)
		if err != nil {
			response.FailWithDetailed("1010","","Token获取失败",c)
			return
		}
		if device.MachineBindState == 3{
			response.FailWithDetailed("1039","","设备已到期",c)
			return
		}

		var data = make(map[string]interface{})
		data["token"] = token
		data["device"] = device
		data["isBind"] = device.MachineBindState
		response.Result("0000",data,"操作成功",c )
	}

}