package routers

import (
	"Shadowsocks-Service-Pack/Panel/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/message", &controllers.BaseController{}, "*:Message")
	beego.Router("/register", &controllers.RegisterController{})

	beego.AutoRouter(&controllers.MailController{})
	beego.AutoRouter(&controllers.RegisterController{})
}
