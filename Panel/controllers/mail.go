package controllers

type MailController struct {
	BaseController
}

func (this *MailController) ActiveMsg() {
	this.Message("请激活", "已发送激活邮件到您的邮箱,请激活帐号.", this.UrlFor("HomeController.Get"), 10)
}
