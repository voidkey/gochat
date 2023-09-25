package models

import (
	"fmt"
	"gochat/utils"

	"gorm.io/gorm"
)

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

func SearchFriends(userId int64) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(v)
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}

//添加好友
// func AddFriend(userId int64, targetId int64) (int, string) {
// 	user := UserBasic{}
// 	if targetId != 0 {
// 		user = FindById(targetId)
// 		if user.Salt != "" {
// 			if userId == user.ID {
// 				return -1,"不可以加自己" //自己不能加自己
// 			}
// 			contact0 := Contact{}
// 			utils.DB.Where("owner_id = ? and target_id = ? and type=1",userId,targetId).Find(&contact0)
// 			if contact0.ID != 0 {
// 				return -1,"不能重复添加好友"
// 			}
// 			tx := utils.DB.Begin()
// 			//事物一旦开始，无论什么异常最终都会rollback
// 			defer func (){
// 				if r := recover(); r!=nil {
// 					tx.Rollback()
// 				}
// 			}()
// 			contact :=  Contact{}
// 			contact.OwnerId = userId
// 			contact.TargetId = targetId
// 			contact.Type = 1
// 			if err := utils.DB.Create(&contact).Error;err!=nil{
// 				tx.Rollback()
// 				return -1,"添加好友失败"
// 			}
// 			contact1 := Contact{}
// 			contact1.OwnerId = targetId
// 			contact1.TargetId = userId
// 			contact1.Type = 1
// 			if err := utils.DB.Create(&contact1).Error;err!=nil{
// 				tx.Rollback()
// 				return -1,"添加好友失败"
// 			}
// 			utils.DB.Create(&contact1)
// 			tx.Commit()
// 			return 0,"添加好友成功"
// 		}
// 		return -1,"没有找到此用户"
// 	}
// 	return -1,"好友ID不能为空"
// }
