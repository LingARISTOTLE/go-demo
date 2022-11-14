package main

import (
	"ginDemo/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//    import _ "github.com/jinzhu/gorm/dialects/mysql"
//    // import _ "github.com/jinzhu/gorm/dialects/postgres"
//    // import _ "github.com/jinzhu/gorm/dialects/sqlite"
//    // import _ "github.com/jinzhu/gorm/dialects/mssql"

func main() {
	//获取数据库连接
	db := common.InitDB()
	//程序执行完出栈执行关闭
	defer db.Close()

	//获取gin引擎
	ginServer := gin.Default()
	//绑定连接路由
	ginServer = CollectRoute(ginServer)
	//抓取异常
	panic(ginServer.Run(":8083"))
}
