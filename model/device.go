package model

import "ginLearn/util"

type Device struct {
	ID uint16 `json:"id,string" form:"id"`
	Num string `json:"num"`
	State int `json:"state" `
	SellerId int `json:"sellerId,string" gorm:"column:currentSellerId" `
	MachineBindState int `json:"machineBindState" gorm:"column:machineBindState" `
	CreateD util.Time `json:"createD" gorm:"column:createD"`
}

