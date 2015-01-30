package controllers

import (
	"Shadowsocks-Service-Pack/Panel/library/mail"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

type BaseController struct {
	beego.Controller
}

var cpt *captcha.Captcha
var IsDev bool
var siteName string
var siteRoot string

func (this *BaseController) Prepare() {
	if beego.RunMode == "dev" {
		IsDev = true
	}

	siteName = beego.AppConfig.String("site::name")
	siteRoot = beego.AppConfig.String("site::root")

	// use beego cache system store the captcha data
	cache := cache.NewFileCache()
	cache.CachePath = "temp/cache"
	cache.FileSuffix = ".cache"
	cpt = captcha.NewWithFilter("/captcha/", cache)

	this.Data["xsrf_token"] = this.XsrfToken()
	this.Data["logout"] = this.UrlFor("LogoutController.Get")
}

type Message struct {
	msg      string
	url      string
	waitTime int
}

var msg Message

/**
 * Jump page
 * author: Frank Kung <kfanjian@gmail.com>
 * date: 2015-01-27 17:28:31
 */
func (this *BaseController) Message(title, msg, url string, wait int) {
	data := make(map[string]interface{})
	data["title"] = title
	data["msg"] = msg
	data["wait"] = wait
	data["url"] = url
	this.Data["msg"] = data
	this.TplNames = "message.tpl"
}

/**
 * Goto Login gateway if not login
 * author: Frank Kung <kfanjian@gmail.com>
 * date: 2015-01-27 15:26:38
 */
func (this *BaseController) needLogin() {
	if this.GetSession("userid") == nil {
		this.Redirect(this.UrlFor("LoginController.Get"), 302)
	}
}

/**
 * Goto Homepage if have login
 * author: Frank Kung <kfanjian@gmail.com>
 * date: 2015-01-27 15:26:46
 */
func (this *BaseController) haveLogin() {
	if this.GetSession("userid") != nil {
		this.Redirect(this.UrlFor("HomeController.Get"), 302)
	}

}

/**
 * Send mail to
 * author: Frank Kung <kanjian@gmail.com>
 * date: 2015-01-30 18:21:45
 */
func (this *BaseController) sendmail(subject, html, toEmail string) bool {
	host := beego.AppConfig.String("tlsmail::host")
	port, _ := beego.AppConfig.Int("tlsmail::port")
	email := beego.AppConfig.String("tlsmail::email")
	password := beego.AppConfig.String("tlsmail::password")
	fromName := beego.AppConfig.String("tlsmail::fromName")
	return mail.Send(host, email, password, toEmail, fromName, subject, html, port)
}
