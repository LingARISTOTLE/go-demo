package dto

import "ginDemo/model"

type UserDto struct { // 只返回给前端名称和手机号
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Username,
		Telephone: user.Telephone,
	}
}
