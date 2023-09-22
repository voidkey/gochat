package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name     string //Sender
	OwenerId uint   //Receiver
	Icon     string //Msg Type
	Type     int
	Desc     string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
