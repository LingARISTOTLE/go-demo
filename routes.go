package main

import (
	"ginDemo/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine { //该方法用于绑定路由
	r.POST("/auto/register", controller.Register)

	return r
}
