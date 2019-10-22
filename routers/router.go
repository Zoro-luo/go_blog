package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
	//注解路由
	beego.Include(&controllers.IndexController{})
	//自动匹配
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.ArticleController{})
}
