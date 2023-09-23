package models

import "gorm.io/gorm"

// Relationship of Users
type Contact struct {
	gorm.Model
	OwnerId  int64 //User
	TargetId int64 //
	Type     int   //
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
