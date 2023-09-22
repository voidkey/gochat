package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId   uint   //Sender
	TargetId uint   //Receiver
	Type     string //Msg Type
	Media    int    //Word Pic Audio
	Content  string //content
	Pic      string
	Url      string
	Desc     string
	Amount   int
}

func (table *Message) TableName() string {
	return "message"
}
