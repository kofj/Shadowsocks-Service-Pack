package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

type BaseController struct {
	beego.Controller
}

var cpt *captcha.Captcha

func (this *BaseController) Prepare() {
	// use beego cache system store the captcha data
	cache := cache.NewFileCache() //cache.NewMemoryCache()
	cache.CachePath = "temp/cache"
	cache.FileSuffix = ".cache"
	cpt = captcha.NewWithFilter("/captcha/", cache)

	this.Data["xsrf_token"] = this.XsrfToken()

}
