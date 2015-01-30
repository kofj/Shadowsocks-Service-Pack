package controllers

import (
	"Shadowsocks-Service-Pack/Panel/library/utils"
	"Shadowsocks-Service-Pack/Panel/models"
	"bytes"
	"fmt"
	"html/template"
)

type MailController struct {
	BaseController
}

type ActiveMail struct {
	Userid     interface{}
	Username   interface{}
	Email      interface{}
	Activelink string
	Sitename   string
	Activecode string
}

func (this *MailController) ActiveMsg() {
	var doc bytes.Buffer
	// get email template from db
	templateValue, _ := models.GetConfig("ActiveMailTemplate")
	// get verify code
	data := fmt.Sprintf("%v%v%v%v%v", this.GetSession("userid"), this.GetSession("username"), this.GetSession("email"), this.GetSession("password"), this.GetSession("invite")) //+  +  +  +
	code := utils.CreateTimeLimitCode(data, 180, nil)

	// set template
	msg := ActiveMail{Sitename: siteName, Userid: this.GetSession("userid"), Username: this.GetSession("username"), Email: this.GetSession("email"), Activecode: code, Activelink: siteRoot + "mail/active?code=" + code}

	// parse template
	t := template.New("template")
	t, tErr := t.Parse(templateValue.Value)
	if tErr != nil {
		fmt.Errorf("[Mail Template Parse Error]%v", tErr)
	}
	t.Execute(&doc, msg)
	subject := fmt.Sprintf("您在%s注册了帐号:%v,请激活", siteName, msg.Username)
	toEmail := fmt.Sprint(msg.Email)
	this.sendmail(subject, doc.String(), toEmail)
	fmt.Println(data, code, templateValue.Value)
	fmt.Println(doc.String())
	this.Message("请激活", "已发送激活邮件到您的邮箱"+toEmail+",请在3小时内激活帐号.", this.UrlFor("HomeController.Get"), 10)
}

func (this *MailController) Name() {
	models.SetConfig("ActiveMailTemplate", "1激活")
}
