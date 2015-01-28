package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	// Need Login.
	this.needLogin()

	this.TplNames = "home.tpl"
}

/**
 * Show user the result of send mail
 * @author Frank Kung <kfanjian@gmail.com>
 * @date 2015-01-28 19:49:11
 */
func SendMail() {

}
