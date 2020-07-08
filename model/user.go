package model

import "ginLearn/util"

type User struct {
	ID uint16 `json:"id,string" form:"id"`
	OpenId string `json:"openId" form:"openId"  gorm:"column:openid"`
	Sex int `json:"sex" `
	NickName string `json:"nickName" form:"nickName"  gorm:"column:nickName"`
	Mobile string `json:"mobile" form:"mobile" `
	HeadImg string `json:"headImg"  from:"headImg"  gorm:"column:headImg"`
	CreateD util.Time `json:"createD" gorm:"column:createD"`
	UpdateD util.Time `json:"updateD" gorm:"column:updateD"`
}

