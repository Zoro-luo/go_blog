package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liteblog/models"
)

//注册界面
// @router /register [get]
func (this *IndexController) Register() {
	this.TplName = "register.html"
}

//注册操作
// @router /register [post]
func (this *IndexController) ResiPost() {
	//1 拿到表单post提交的数据
	userName := this.GetString("userName")
	pwd := this.GetString("pwd")
	//2 数据校验
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		//跳转重定向
		this.Redirect("/register", 302)
		return
	}
	//3 插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据库失败", err)
		this.Redirect("/register", 302)
		return
	}
	//4 注册成功跳转登录界面
	this.Redirect("/login", 302) //跳转 速度快
	//this.TplName = "login.html"				//渲染模板 可以传递数据到模板
	//this.Ctx.WriteString("注册成功!")			//浏览器端直接输出
}

//登录界面
// @router /login [get]
func (this *IndexController) Login() {
	this.TplName = "login.html"
}

//登录操作
// @router /login [post]
func (this *IndexController) HandleLogin() {
	//1 拿到数据
	userName := this.GetString("userName")
	pwd := this.GetString("pwd")
	//2 判断数据是否合法
	if userName == "" || pwd == "" {
		beego.Info("用户名或密码不能为空")
		this.TplName = "login.html"
		return
	}
	//3 查询数据是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("查询失败")
		this.TplName = "login.html"
		return
	}
	//4 成功跳转
	this.Redirect("/login", 302)
	//this.Ctx.WriteString("欢迎你,登录成功!")
}
