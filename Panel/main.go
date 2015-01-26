package main

import (
	_ "Shadowsocks-Service-Pack/Panel/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true
	beego.SessionProvider = "file"
	beego.SessionGCMaxLifetime = 36000 // 10 hours
	beego.SessionName = "xssid"
	beego.SessionCookieLifeTime = 36000 // 10 hours
	beego.SessionAutoSetCookie = true
	beego.SessionSavePath = "temp/session"

	beego.Run()
}
