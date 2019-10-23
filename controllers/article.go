package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liteblog/models"
	"math"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	BaseController
}

//文章列表页
func (c *ArticleController) List() {
	o := orm.NewOrm()
	var articles []models.Article
	res := o.QueryTable("tb_article")
	pageIndex := c.GetString("pageIndex")
	pageCurrent,err := strconv.Atoi(pageIndex)
	if err != nil {
		pageCurrent = 1		//当前页
	}
	//_, err := res.All(&articles)
	count, _ := res.Count()	//总条数
	pageSize := 2		//每页显示条数
	start := pageSize * (pageCurrent - 1)
	//分页查询
	_,err = res.Limit(pageSize, start).All(&articles)
	//总页数
	pageCount := float64(count) / float64(pageSize)
	pageCount = math.Ceil(pageCount) 	//向上取整
	if err != nil {
		beego.Info("查询所有文章信息出错")
		return
	}
	if pageCurrent < 1 {
		pageCurrent = 1
	}
	if pageCurrent > int(pageCount) {
		pageCurrent = int(pageCount)
	}
	c.Data["pageCurrent"] = pageCurrent
	c.Data["pageCount"] = pageCount
	c.Data["count"] = count
	c.Data["articles"] = articles
	c.TplName = "admin/article-list.html"
}

//文章详情页
func (c *ArticleController) Detail() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("找不到指定的文章Id", err)
		return
	}
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	c.Data["article"] = arti
	c.TplName = "admin/article-detail.html"
}

//文章添加
func (c *ArticleController) Add() {
	if c.Ctx.Request.Method == "POST" {
		//post表单提交
		articleName := c.GetString("articleName")
		articleContent := c.GetString("articleContent")
		f, h, err := c.GetFile("articleFile")
		defer f.Close()
		//限定文件格式
		fileExt := path.Ext(h.Filename)
		if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg" {
			beego.Info("上传文件格式错误")
			return
		}
		//限定文件大小
		if h.Size > 500000000 {
			beego.Info("上传文件过大")
			return
		}
		//限定文件名重复
		fileName := time.Now().Format("15:04:05") + fileExt
		if err != nil {
			beego.Info("上传文件失败", err)
		} else {
			//文件保存
			beego.Info("上传文件成功")
			c.SaveToFile("articleFile", "/static/img/"+fileName)
		}
		//数据校验
		if articleName == "" || articleContent == "" {
			beego.Info("添加文章数据错误", err)
			return
		}
		o := orm.NewOrm()
		arti := models.Article{}
		arti.Aname = articleName
		arti.Acontent = articleContent
		arti.Aimg = "/static/img/" + fileName
		_, err = o.Insert(&arti)
		if err != nil {
			beego.Info("插入数据库错误")
			return
		}
		beego.Info("插入数据成功")
		//插入成功跳转到文章列表页
		c.Redirect("/article/list", 302)
	}
	c.TplName = "admin/article-add.html"
}

//文章编辑
func (c *ArticleController) Update() {
	if c.Ctx.Request.Method == "POST" {
		id, _ := c.GetInt("id")
		beego.Info("id is ", id)
		articleName := c.GetString("articleName")
		articleContent := c.GetString("articleContent")
		f, h, err := c.GetFile("articleFile")
		defer f.Close()
		fileExt := path.Ext(h.Filename)
		if fileExt != ".jpg" && fileExt != ".png" {
			beego.Info("上传文件格式错误")
			return
		}
		if h.Size > 50000000 {
			beego.Info("上传文件过大")
			return
		}
		fileName := time.Now().Format("2016-01-02 15:04:05") + fileExt
		if err != nil {
			beego.Info("上传文件失败")
			return
		} else {
			c.SaveToFile("articleFile", "./static/img/"+fileName)
		}
		if articleName == "" || articleContent == "" {
			beego.Info("更新数据失败")
			return
		}
		o := orm.NewOrm()
		arti := models.Article{Id: id}
		//更新要先根据id先查询
		err = o.Read(&arti)
		if err != nil {
			beego.Info("查询数据错误")
			return
		}
		//赋值更新操作
		arti.Aname = articleName
		arti.Acontent = articleContent
		arti.Aimg = "./static/img/" + fileName
		_, err = o.Update(&arti, "Aname", "Acontent", "Aimg")
		if err != nil {
			beego.Info("更新数据错误")
			return
		}
		c.Redirect("/article/list", 302)

	}
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("找不到指定的文章id", err)
		return
	}
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	c.Data["article"] = arti
	c.TplName = "admin/article-update.html"
}

//文章删除
func (c *ArticleController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("文章ID获取失败", err)
		return
	}
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}
	_, err = o.Delete(&arti)
	if err != nil {
		beego.Info("删除失败", err)
		return
	}
	c.Redirect("/article/list", 302)
}
