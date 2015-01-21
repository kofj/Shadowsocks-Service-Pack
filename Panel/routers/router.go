package routers

import (
	"Shadowsocks-Service-Pack/Panel/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
