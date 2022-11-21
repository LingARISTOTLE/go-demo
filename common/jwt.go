package common

import (
	"ginDemo/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) { //生成token
	//设置token过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}

	//使用jwt密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) { //解析token
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
