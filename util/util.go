package util

import (
	"math/rand"
	"time"
)

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
