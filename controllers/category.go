package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liteblog/models"
)

type CateController struct {
	BaseController
}

//分类列表页
func (c *CateController) List() {
	o := orm.NewOrm()
	var cates []models.Category
	_, err := o.QueryTable("go_category").All(&cates)
	if err != nil {
		beego.Info("查询分类数据错误")
	}
	c.Data["cates"] = cates
	c.TplName = "admin/cate-list.html"
}

//添加分类
func (c *CateController) Add() {
	if c.Ctx.Request.Method == "POST" {
		cateName := c.GetString("cateName")
		if cateName == "" {
			beego.Info("不能添加空的数据类型")
			return
		}
		o := orm.NewOrm()
		cate := models.Category{}
		cate.TypeName = cateName
		_, err := o.Insert(&cate)
		if err != nil {
			beego.Info("分类数据插入失败")
			return
		}
		c.Redirect("/cate/list", 302)
	}
	c.TplName = "admin/cate-add.html"
}
