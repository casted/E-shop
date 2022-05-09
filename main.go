package main

import (
	"encoding/gob"

	"LeastMall/common"
	"LeastMall/models"
	_ "LeastMall/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	//添加方法到map,用于前端html调用
	beego.AddFuncMap("timestampToDate", common.TimestampToDate)
	models.DB.LogMode(true)
	beego.AddFuncMap("formatImage", common.FormatImage)
	beego.AddFuncMap("mul", common.Mul)
	beego.AddFuncMap("formatAttribute", common.FormatAttribute)
	beego.AddFuncMap("setting", models.GetSettingByColumn)

	//后台配置允许跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"127.0.0.1"},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type"},
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type"},
		AllowCredentials: true, //是否允许cookie
	}))
	//注册模型
	gob.Register(models.Administrator{})
	//关闭数据库
	//defer models.DB.Close()
	//配置redis用于存储session
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//docker-compose 请设置为redisServiceHost
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "redisServiceHost:6379"

	//本地启动，请设置如下
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.Run()
}
