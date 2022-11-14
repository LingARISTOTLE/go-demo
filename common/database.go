package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

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

	//给属性赋值
	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}
