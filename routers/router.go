package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
	//注解路由
	beego.Include(&controllers.IndexController{}) 			//前台
	//自动匹配
	beego.AutoRouter(&controllers.AdminController{}) 		//后台用户
	beego.AutoRouter(&controllers.ArticleController{})		//文章
	beego.AutoRouter(&controllers.CateController{})			//分类
}
