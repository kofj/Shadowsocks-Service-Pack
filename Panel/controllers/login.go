package controllers

import (
	. "Shadowsocks-Service-Pack/Panel/models"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	// Is Login?
	this.haveLogin()

	userid := this.GetSession("userid")
	if userid != nil {
		this.Redirect(this.UrlFor("HomeController.Get"), 302)
	} else {
		this.TplNames = "login.tpl"
	}
}

func (this *LoginController) Post() {

	username := this.GetString("username")
	password := this.GetString("password")
	captcha := cpt.VerifyReq(this.Ctx.Request)

	if !captcha { // Check captcha.
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid captcha", "refer": "/"}
		this.ServeJson()
		return
	}
	if username == "" || password == "" { // Check input
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
		this.ServeJson()
	}

	// Get data.
	user, err := FindUser(username)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "user does not exist", "refer": "/"}
	} else {
		if password == user.Password {

			this.SetSession("userid", user.Id)
			this.SetSession("username", user.Username)
			this.SetSession("status", user.Status)

			this.Data["json"] = map[string]interface{}{"result": true, "msg": "user[" + user.Username + "] login success ", "refer": this.UrlFor("HomeController.Get")}
		} else {
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "login failed ", "refer": "/"}
		}
	}
	this.ServeJson()
}
