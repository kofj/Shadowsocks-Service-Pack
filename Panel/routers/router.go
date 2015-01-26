package routers

import (
	"Shadowsocks-Service-Pack/Panel/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/Home", &controllers.HomeController{})
}
