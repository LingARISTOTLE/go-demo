package model

type User struct {
	//gorm.Model
	Id        int    //id
	Username  string //用户名
	Telephone string //电话
	Password  string //密码
}
