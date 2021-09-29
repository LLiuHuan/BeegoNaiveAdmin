package main

import (
	_ "BeegoNaiveAdmin/initialize"
	_ "BeegoNaiveAdmin/routers"

	"github.com/beego/beego/v2/server/web/filter/cors"

	_ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/client/orm"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//InsertFilter是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	beego.Run()
}
