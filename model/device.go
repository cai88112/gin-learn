package model

import "ginLearn/util"

type Device struct {
	ID uint16 `json:"id,string" form:"id"`
	Num string `json:"num"`
	State int `json:"state" `
	SellerId int `json:"sellerId,string" gorm:"column:currentSellerId" form:"sellerId"`
	Name string `json:"name"`
	Pwd string `json:"pwd"`
	CreateD util.Time `json:"createD" gorm:"column:createD"`
}

