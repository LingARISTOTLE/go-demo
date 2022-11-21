package main

import (
	"ginDemo/controller"
	"ginDemo/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine { //该方法用于绑定路由
	//注册事件
	r.POST("/auto/register", controller.Register)

	//登录事件
	r.POST("/auto/login", controller.Login)

	//用户信息
	r.GET("/auto/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
