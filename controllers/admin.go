package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liteblog/models"
)

type AdminController struct {
	BaseController
}

//登陆
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		if username == "" || password == "" {
			beego.Info("用户名或密码不能为空")
			c.Redirect("/admin/login", 302)
			return
		}
		o := orm.NewOrm()
		admin := models.Admin{UserName: username} //查询对象的条件
		err := o.Read(&admin, "UserName")         //根据UserName查询
		if err != nil {
			beego.Info("查询失败!", err)
			c.TplName = "admin/login.html"
			return
		}
		c.SetSession("username", username)
		c.Redirect("index.html", 302)
		//c.Ctx.WriteString("登陆成功！！")
	}

	c.TplName = "admin/login.html"
}

//后台首页
func (c *AdminController) Index() {
	c.TplName = "admin/index.html"
}

//后台注册了一个登陆管理员
func (c *AdminController) Register() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		if username == "" || password == "" {
			beego.Info("用户名或密码不能为空")
			return
		}
		o := orm.NewOrm()
		admin := models.Admin{UserName: username}
		admin.UserName = username
		admin.Password = password
		_, err := o.Insert(&admin)
		if err != nil {
			beego.Info("插入数据库失败", err)
			c.Redirect("/admin/register", 302)
			return
		}
		c.Ctx.WriteString("注册成功")
	}
	c.TplName = "admin/register.html"
}
