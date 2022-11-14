package controller

import (
	"ginDemo/common"
	"ginDemo/model"
	"ginDemo/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(context *gin.Context) {
	db := common.GetDB()

	//获取参数
	username := context.PostForm("username")
	password := context.PostForm("password")
	telephone := context.PostForm("telephone")

	//数据认证
	if len(telephone) != 11 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号不是11位"})
	}

	//数据认证
	if len(password) < 6 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
	}

	//判断手机号是否存在
	//...实现
	log.Println(username, password, telephone)
	if isTelephoneExist(db, telephone) {
		//如果用户存在就不允许注册
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该电话对应的用户已存在"})
		return
	}

	//判断名称，如果没有自动生成
	if len(username) == 0 {
		//自动生成10为随机字符串的用户名
		username = util.GetRandomString(10)
	}

	log.Println(username, password, telephone)

	//创建用户
	newUser := model.User{
		Username:  username,
		Telephone: telephone,
		Password:  password,
	}

	//将用户注册进数据库
	db.Create(&newUser)

	context.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool { //查询手机号
	var user model.User
	//查询，将查询到的第一个结果封装到user对象中
	db.Where("telephone = ?", telephone).First(&user)

	if user.Id != 0 {
		return true
	}
	return false
}
