package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	// Need Login.
	this.needLogin()

	this.TplNames = "home.tpl"
}
