package model

type User struct {
	//gorm.Model
	Id        uint   //id
	Username  string //用户名
	Telephone string //电话
	Password  string //密码
}
