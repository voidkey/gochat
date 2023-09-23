package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name     string //Sender
	OwenerId int64  //Receiver
	Icon     string //Msg Type
	Type     int
	Desc     string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
