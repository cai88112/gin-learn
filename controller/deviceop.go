package controller

import (
	"crypto/md5"
	"fmt"
	"ginLearn/common"
	"ginLearn/global"
	"ginLearn/log"
	"ginLearn/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)
var mylog = log.GetLogger("device.log")

type clientManager struct {
	State int
	UserName string `gorm:"column:userName"`
	Pwd string
	RealName string `gorm:"column:realName"`
}

func DeviceLogin(c *gin.Context)  {
	username := c.PostForm("userName")
	pwd := c.PostForm("pwd")
	num := c.PostForm("deviceNum")
	var deviceModel model.Device
	 var result = clientManager{}
	manager :=global.Db.Table("client_manager").Where("userName=?",username).Scan(&result).RecordNotFound()
	mylog.Info(result)
	//方法一
	data := []byte(pwd)
	has := md5.Sum(data)
	md5pwd := fmt.Sprintf("%x", has) //将[]byte转成16进制
	if manager{
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "账户或密码错误",
		})
		return
	}else if !strings.EqualFold(result.Pwd,md5pwd){
		c.JSON(401, gin.H{
			"code": 441,
			"msg":  "账户或密码错误",
		})
		return
	}
	//flag :=db.Raw("SELECT id,num,state,currentSellerId sellerId FROM oil_client WHERE num = ?", num).Scan(&deviceModel).RecordNotFound()
	flag := global.Db.Table("oil_client").Where("num = ?",num).First(&deviceModel).RecordNotFound()

	if flag{
		mylog.WithFields(logrus.Fields{
			"deviceNum":num,
			"userName":username,
			"pwd":pwd,
		}).Error("无法找不到设备")
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "device not found",
		})
	}else {
		device := model.Device{
			Num:      num,
			State:    3,
			SellerId: 160,
			Name:     username,
			Pwd:      pwd,
		}
		token, err := common.GetToken(device)
		if err != nil {
			fmt.Printf("token create fail %s", err.Error())
			c.JSON(400, gin.H{
				"code": 401,
				"msg":  "token create fail",
			})
			return
		}

		c.JSON(200, gin.H{
			"token": token,
			"device":deviceModel,
		})
	}

}