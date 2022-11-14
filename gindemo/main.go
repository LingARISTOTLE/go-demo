package test1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义拦截器,go语言的中间件
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("user", "userid - 1")

		if context.Query("username") == "zhangsan" {
			//放行
			context.Next()
		} else {
			//组织
			context.Abort()
		}
	}
}

func main() {
	ginServer := gin.Default()

	ginServer.Use(myHandler())

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "HelloWorld,GET"})
	})

	//gin restful风格
	ginServer.POST("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "HelloWorld,POST"})
	})

	ginServer.PUT("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "HelloWorld,PUT"})
	})

	ginServer.DELETE("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "HelloWorld,DELETE"})
	})

	getInfo(ginServer)

	getRestful(ginServer)

	getJson(ginServer)

	getForm(ginServer)

	testRedirect(ginServer)

	testMiddle(ginServer)

	ginServer.Run(":8082")
}

func getInfo(ginServer *gin.Engine) {
	ginServer.GET("/hello/info", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		context.JSON(200, gin.H{"username": username, "password": password})
	})
}

func getRestful(ginServer *gin.Engine) {
	ginServer.GET("/hello/info/:username/:password", func(context *gin.Context) {
		username := context.Param("username")
		password := context.Param("password")
		context.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})
}

func getJson(ginServer *gin.Engine) {
	ginServer.POST("/hello/json", func(context *gin.Context) {
		result, _ := context.GetRawData()
		//设置key为string ，value为interface任意类
		var m map[string]interface{}
		//将json解析到map中
		_ = json.Unmarshal(result, &m)
		context.JSON(http.StatusOK, m)
	})
}

func getForm(ginServer *gin.Engine) {
	ginServer.POST("/hello/form", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})
}

func testRedirect(ginServer *gin.Engine) {
	ginServer.GET("/toBaidu", func(context *gin.Context) {
		context.Redirect(301, "https://www.baidu.com/")
	})
}

func testMiddle(ginServer *gin.Engine) {
	ginServer.GET("/testMiddle", myHandler(), func(context *gin.Context) {
		var user string = context.MustGet("user").(string)
		context.JSON(http.StatusOK, gin.H{"user": user})
	})
}
