package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"ginLearn/global"
	"ginLearn/model"
	"strings"
)



type clientManager struct {
	State int
	UserName string `gorm:"column:userName"`
	Pwd string
	RealName string `gorm:"column:realName"`
}

func ValidAccount(userName,pwd  string) error{
	var result = clientManager{}
	err :=global.Db.Table("client_manager").Where("userName=?",userName).Scan(&result).Error
	if err != nil{
		return err
	}
	data := []byte(pwd)
	has := md5.Sum(data)
	md5pwd := fmt.Sprintf("%x", has) //将[]byte转成16进制
	if !strings.EqualFold(md5pwd,result.Pwd){
		return errors.New("账户或密码错误")
	}
	return nil
}

func ValidDevice(deviceNum string) (model.Device,error){
	var deviceModel model.Device
	err := global.Db.Table("oil_client").Where("num = ?",deviceNum).First(&deviceModel).Error

	if err != nil{
		return deviceModel,err
	}
	return deviceModel,nil

}