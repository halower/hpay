package main

import (
	"github.com/halower/hipay/databse"
	_ "github.com/halower/hipay/docs"
	"github.com/halower/hipay/router"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title 易支付
// @version 1.0
// @description 个人点对点支付服务
// @contact.name halower
// @contact.url https://github.com/halower
// @license.name 源码协议
// @license.url https://github.com/halower/newbie-spring-boot-project/blob/master/license_996.txt
func main() {
	defer databse.SqlDB.Close()
	r := router.InitRouter()
	r.Static("/pay", "./public")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":80")
}