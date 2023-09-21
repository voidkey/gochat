package models

import (
	"fmt"
	"gochat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LogInTime     time.Time `gorm:"column:log_in_time" json:"log_in_time"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time" json:"heartbeat_time"`
	LogOutTime    time.Time `gorm:"column:log_out_time" json:"log_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
