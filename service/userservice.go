package service

import (
	"fmt"
	"gochat/models"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已注册！",
		})
		return
	}

	if password != repassword {
		c.JSON(4000327, gin.H{
			"message": "两次密码不一致！",
		})
	} else {
		user.Password = password
	}
	user.LogInTime = time.Now()
	user.HeartbeatTime = time.Now()
	user.LogOutTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "添加用户成功！",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "ID"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功！",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "ID"
// @param name formData string false "Username"
// @param password formData string false "Password"
// @param phone formData string false "Phone"
// @param email formData string false "Email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println("govalidator err: ", err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配！",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "修改用户成功！",
		})
	}

}
