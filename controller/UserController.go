package controller

import (
	"ginDemo/common"
	"ginDemo/dto"
	"ginDemo/model"
	"ginDemo/response"
	"ginDemo/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
		return
	}

	//数据认证
	if len(password) < 6 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
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
	//首先修改密码，我们需要对用户密码进行加密
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(500, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Username:  username,
		Telephone: telephone,
		Password:  string(hasPassword),
	}

	//将用户注册进数据库
	db.Create(&newUser)

	//context.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"msg":  "注册成功",
	//})

	response.Success(context, nil, "注册成功")
}

func Login(context *gin.Context) {
	db := common.GetDB()
	//获取参数
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")

	//数据认证
	if len(telephone) != 11 {
		//更新一下新的响应格式
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号不是11位")
		//context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号不是11位"})
		return
	}

	//数据认证
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		//context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.Id == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		//context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//此时数据库中的密码是加密后的密码
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(context, http.StatusBadRequest, 400, nil, "密码错误")
		//context.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//发放token
	//token := "11" //目前先写死数据
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")
		//context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		return
	}

	//返回结果
	//context.JSON(200, gin.H{
	//	"code": 200,
	//	"data": gin.H{"token": token},
	//	"msg":  "登录成功",
	//})
	response.Success(context, gin.H{"token": token}, "登录成功")
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

func Info(ctx *gin.Context) { //获取用户信息
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
}
