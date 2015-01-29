package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	// Need Login.
	this.needLogin()

	if status := this.GetSession("status"); status == 0 {
		this.TplNames = "notactive.tpl"
	} else {
		this.TplNames = "home.tpl"
	}
}
