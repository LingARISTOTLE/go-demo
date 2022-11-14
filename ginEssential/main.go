package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//    import _ "github.com/jinzhu/gorm/dialects/mysql"
//    // import _ "github.com/jinzhu/gorm/dialects/postgres"
//    // import _ "github.com/jinzhu/gorm/dialects/sqlite"
//    // import _ "github.com/jinzhu/gorm/dialects/mssql"

type User struct {
	//gorm.Model
	Id        int    //id
	Username  string //用户名
	Telephone string //电话
	Password  string //密码
}

func main() {
	//获取gin引擎
	ginServer := gin.Default()

	//获取数据库连接
	db := InitDB()
	//程序执行完出栈执行关闭
	defer db.Close()

	//注册
	ginServer.POST("/auto/register", func(context *gin.Context) {
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
			username = GetRandomString(10)
		}

		log.Println(username, password, telephone)

		//创建用户
		newUser := User{
			Username:  username,
			Telephone: telephone,
			Password:  password,
		}

		//将用户注册进数据库
		db.Create(&newUser)

		context.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})

	})

	panic(ginServer.Run(":8083"))
}

func isTelephoneExist(db *gorm.DB, telephone string) bool { //查询手机号
	var user User
	//查询，将查询到的第一个结果封装到user对象中
	db.Where("telephone = ?", telephone).First(&user)

	if user.Id != 0 {
		return true
	}
	return false
}

func GetRandomString(length int) string { //自动生成字符串
	//生成随机数
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	//创建新的无参数组
	result := make([]byte, length)

	//获取unix时间戳作为随机数
	rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func InitDB() *gorm.DB { //初始化数据库连接词
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gintest"
	username := "root"
	password := "239732"
	charset := "utf8"
	//连接args和java的类似，一条成形的如下:
	//"root:root123@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("数据库连接失败，异常err:" + err.Error())
	}
	return db
}
