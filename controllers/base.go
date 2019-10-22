package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	//拿到当前访问的路由
	this.Data["Path"] = this.Ctx.Request.RequestURI
}
