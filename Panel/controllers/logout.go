package controllers

type LogoutController struct {
	BaseController
}

func (this *LogoutController) Get() {
	this.DestroySession()

	// Jump page
	this.Message("成功注销", "您已经成功退出系统", this.UrlFor("LoginController.Get"), 3)
}
