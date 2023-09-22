package main

import (
	"fmt"
	"gochat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:271828@tcp(127.0.0.1:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// 迁移 schema
	//db.AutoMigrate(&models.UserBasic{})
	//db.AutoMigrate(&models.Message{})
	//db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.GroupBasic{})

	// // Create
	// user := &models.UserBasic{}
	// user.Name = "飛"
	// db.Create(user)

	// // Read
	// db.First(user, 1) // 根据整型主键查找
	// //db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// // Update - 将 user 的 price 更新为 200
	// db.Model(user).Update("Password", "114514")
	// // Update - 更新多个字段
	// db.Model(user).Updates(models.UserBasic{Phone: "110", Email: "a@b.com"}) // 仅更新非零值字段
	// db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 user
	// db.Delete(user, 1)
}
