package main

import (
	"ginDemo/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

//    import _ "github.com/jinzhu/gorm/dialects/mysql"
//    // import _ "github.com/jinzhu/gorm/dialects/postgres"
//    // import _ "github.com/jinzhu/gorm/dialects/sqlite"
//    // import _ "github.com/jinzhu/gorm/dialects/mssql"

//import (
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//	"github.com/spf13/viper"
//	"jkdev.cn/api/common"
//)

func main() {
	//项目启动立马读取配置文件
	InitConfig()
	//获取数据库连接
	db := common.InitDB()
	//程序执行完出栈执行关闭
	defer db.Close()

	//获取gin引擎
	ginServer := gin.Default()
	//绑定连接路由
	ginServer = CollectRoute(ginServer)

	port := viper.GetString("server.port")
	if port != "" {
		port = ":" + port
	} else {
		port = ":8080"
	}
	//抓取异常
	panic(ginServer.Run(port))
}

func InitConfig() { //初始化配置文件
	workDir, _ := os.Getwd()
	//设置配置文件名称
	viper.SetConfigName("application")
	//设置文件类型
	viper.SetConfigType("yml")
	//设置文件目录
	viper.AddConfigPath(workDir + "/config")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
