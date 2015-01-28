package controllers

import (
	. "Shadowsocks-Service-Pack/Panel/models"
	"fmt"
	"regexp"
	"time"
)

type RegisterController struct {
	BaseController
}

/**
 * Display register page.
 * @author Frank Kung <kfanjian@gmail.com>
 * @date 2015-01-28 10:55:34
 */
func (this *RegisterController) Get() {
	this.haveLogin()
	this.Data["invite"] = this.GetString("invite")
	this.TplNames = "register.tpl"
}

/**
 * register post.
 * @author Frank Kung <kfanjian@gmail.com>
 * @date 2015-01-28 10:55:34
 */
func (this *RegisterController) Post() {
	this.haveLogin()

	captcha := cpt.VerifyReq(this.Ctx.Request)

	if !captcha { // Check captcha.
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid captcha"}
		this.ServeJson()
		return
	}

	var user User

	user.Invite = this.GetString("invite")
	user.Username = this.GetString("username")
	user.Password = this.GetString("password")
	user.Email = this.GetString("email")
	user.Createtime = time.Now().Unix()
	user, err := AddUser(user)
	if err != nil {
		if IsDev {
			fmt.Println(err)
		}
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "error"}
		this.ServeJson()
	}
	if err == nil && user.Id > 0 {
		this.SetSession("userid", user.Id)
		this.SetSession("email", user.Email)
		this.SetSession("username", user.Username)
		this.SetSession("status", user.Status)

		this.Data["json"] = map[string]interface{}{"result": true, "msg": "Resgister sucess.", "refer": this.UrlFor("MailController.ActiveMsg")}
		this.ServeJson()
	}
}

/**
 * Check Username
 * @author Frank Kung <kfanjian@gmail.com>
 * @date 2015-01-28 11:19:32
 */
func (this *RegisterController) CheckUsername() {
	value := this.Ctx.Input.Params["0"]

	// check length
	if len(value) < 4 || len(value) > 20 {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "Invalid length."}
		this.ServeJson()
	}

	// Get data.
	user, err := FindUser(value)

	//fmt.Println(user, err)

	if err == nil && user.Id == 0 {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "This username can be register"}
	}
	if err == nil && user.Id != 0 {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "This username cann't be register"}
	}
	this.ServeJson()
	return
}

/**
 * Check Invite Code
 * @author Frank Kung <kfanjian@gmail.com>
 * @date  2015-01-28 11:21:16
 */
func (this *RegisterController) CheckInvite() {
	value := this.Ctx.Input.Params["0"]

	// check length
	if len(value) != 9 {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "Invalid length."}
		this.ServeJson()
	}
	// Get data.
	code, err := FindCode(value)
	//fmt.Println(code, err)

	if err == nil && code.Id != 0 && code.Used == 0 {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "This code can be register"}
	}
	if err == nil && code.Id != 0 && code.Used == 1 {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "This code cann't be register"}
	}
	this.ServeJson()
}

/**
 * Check Email address.
 * @author Frank Kung <kfanjian@gmail.com>
 * @date 2015-01-28 12:04:59
 */
func (this *RegisterController) CheckEmail() {
	value := this.Ctx.Input.Params["0"]

	reg := regexp.MustCompile("^\\w+@\\w+\\.(asia|biz|cat|com|coop|edu|gov|int|info|jobs|mil|mobi|museum|name|net|org|pro|tel|travel|xxx|al|dz|af|ar|ae|aw|om|za|eg|et|ie|ee|ad|ao|ai|ag|at|au|mo|bb|pg|bs|pk|py|ps|bh|pa|br|by|bm|bg|mp|bj|be|is|pr|ba|pl|bo|bz|bw|bt|bf|bi|bv|kp|gq|dk|de|tl|tp|tg|dm|do|ru|ec|er|fr|fo|pf|gf|tf|va|ph|fj|fl|cv|fk|gm|cg|cd|co|cr|gg|gd|gl|ge|cu|gp|gu|gy|kz|ht|kr|nl|an|hm|me|hn|hk|ki|dj|kg|gn|gw|ca|gh|ga|kh|cz|zw|cm|qa|ky|km|ci|kw|cc|hr|ke|ck|lv|ls|la|lb|lt|lr|ly|li|re|lu|rw|ro|mg|im|mv|mt|mw|my|ml|mk|mh|mq|yt|mu|mr|us|um|as|vi|mn|ms|bd|pe|fm|mm|md|ma|mc|mz|mx|mo|nr|np|ni|ne|ng|nu|no|nf|nf|na|za|aq|gs|ss|eu|pw|pn|pt|jp|se|ch|sv|ws|yu|sl|sn|cy|sc|sa|cx|st|sh|kn|lc|sm|pm|vc|lk|sk|si|sj|sz|sd|sr|sb|so|tj|tw|th|tz|to|tc|tt|tn|tv|tr|tm|tk|wf|vu|gt|ve|bn|ug|ua|uy|uz|es|eh|gr|hk|sg|nc|nz|hu|sy|jm|am|ac|ye|iq|ir|il|in|id|uk|vg|io|jo|vn|zm|je|td|gi|cl|cf|cn|yr|yu)$")
	if reg.MatchString(value) == false {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid email"}
		this.ServeJson()
	}

	// Get data.
	user, err := FindUserByMail(value)
	//fmt.Println(user, err)

	if err == nil && user.Id == 0 {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": "This email can be register"}
	}
	if err == nil && user.Id != 0 {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "This email cann't be register"}
	}
	this.ServeJson()
}
