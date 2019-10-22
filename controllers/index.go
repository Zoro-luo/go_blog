package controllers

//渲染首页各界面 [orm增删改查操作]
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liteblog/models"
)

type IndexController struct {
	BaseController
}

//首页界面 [orm插入案例]
// @router / [get]
func (this *IndexController) Index() {
	// 插入:
	o := orm.NewOrm()     //1 orm对象
	user := models.User{} //2 有一个要插入数据的数据库结构体对象
	user.Name = "tom"     //3 对象结构体赋值
	user.Pwd = "123456"
	_, err := o.Insert(&user) //4 插入
	if err != nil {
		beego.Info("插入失败", err)
		return
	}
	this.TplName = "home/index.html"
}

//留言视图 [orm查询案例]
// @router /message [get]
func (this *IndexController) IndexMessage() {
	//查询：
	o := orm.NewOrm()
	user := models.User{}
	user.Name = "tom"            //查询对象的条件
	err := o.Read(&user, "Name") //根据Name查数据
	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	beego.Info("查询成功", user)
	this.TplName = "home/message.html"
}

//关于视图 [orm更新案例]
// @router /about [get]
func (this *IndexController) IndexAbout() {
	//更新：
	o := orm.NewOrm()
	user := models.User{}
	user.Name = "tom"
	err := o.Read(&user, "Name")
	if err == nil { //没有范围错 则查找成功
		user.Name = "jack"
		user.Pwd = "111111"
		_, err := o.Update(&user) //4 更新
		if err != nil {
			beego.Info("更新失败", err)
			return
		}
	}
	this.TplName = "home/about.html"
}

//评论视图 [orm删除案例]
// @router /comment [get]
func (this *IndexController) IndexComment() {
	// 删除:
	o := orm.NewOrm()
	user := models.User{}
	user.Id = 1
	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("删除失败", err)
		return
	}
	this.TplName = "home/comment.html"
}
