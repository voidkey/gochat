package models

import "gorm.io/gorm"

// Relationship of Users
type Contact struct {
	gorm.Model
	OwnerId  uint //User
	TargetId uint //
	Type     int  //
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
