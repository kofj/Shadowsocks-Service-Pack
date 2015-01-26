package controllers

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	this.TplNames = "login.tpl"
}

func (this *LoginController) Post() {
	//
	//this.Ctx.WriteString("Post")
	this.Redirect(this.UrlFor("HomeController.Get"), 302)
}
