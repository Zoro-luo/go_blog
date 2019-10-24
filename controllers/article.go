package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"liteblog/models"
	"liteblog/util"
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
	if c.Ctx.Request.Method == "POST" {
		cateId := c.GetString("category")
		if cateId == "" {
			beego.Info("下拉框传递数据失败")
			return
		}
		beego.Info(cateId)
		//o := orm.NewOrm()
		//var artiCates []models.Article
		//o.QueryTable("tb_article").RelatedSel("Category").Filter("Category__TypeName",cateId).All(&artiCates)
		//beego.Info(artiCates)
	}
	o := orm.NewOrm()
	var articles []models.Article
	res := o.QueryTable("go_article")
	pageIndex := c.GetString("pageIndex")
	pageCurrent, err := strconv.Atoi(pageIndex)
	if err != nil {
		pageCurrent = 1 //当前页
	}
	count, _ := res.Count() //总条数
	pageSize := 3           //每页显示条数
	start := pageSize * (pageCurrent - 1)
	//分页查询
	_, err = res.Limit(pageSize, start).All(&articles)
	//总页数
	pageCount := float64(count) / float64(pageSize)
	pageCount = math.Ceil(pageCount) //向上取整
	util.CheckErr(err,"查询所有文章信息出错")
	FirstPage := false //标识是否首页
	EndPage := false   //标识是否末页
	if pageCurrent == 1 {
		FirstPage = true
	}
	if pageCurrent == int(pageCount) {
		EndPage = true
	}
	//获取分类数据
	var cates []models.Category
	o.QueryTable("go_category").All(&cates)
	c.Data["cates"] = cates
	c.Data["FirstPage"] = FirstPage
	c.Data["EndPage"] = EndPage
	c.Data["pageCurrent"] = pageCurrent
	c.Data["pageCount"] = pageCount
	c.Data["count"] = count
	c.Data["articles"] = articles
	c.TplName = "admin/article-list.html"
}

//文章详情页
func (c *ArticleController) Detail() {
	id, err := c.GetInt("id")
	util.CheckErr(err, "找不到指定的文章Id")
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	util.CheckErr(err, "查询失败")
	c.Data["article"] = arti
	c.TplName = "admin/article-detail.html"
}

//文章添加
func (c *ArticleController) Add() {
	if c.Ctx.Request.Method == "POST" { //post表单提交
		articleName := c.GetString("articleName")
		articleContent := c.GetString("articleContent")
		articleCateId := c.GetString("articleCate")
		beego.Info("文章分类id= ", articleCateId)
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
		arti.Acateid, _ = strconv.Atoi(articleCateId)
		arti.Aimg = "/static/img/" + fileName
		_, err = o.Insert(&arti)
		util.CheckErr(err, "插入数据库错误")
		c.Redirect("/article/list", 302)
	} else { //get提交
		var cates []models.Category
		o := orm.NewOrm()
		_, err := o.QueryTable("go_category").All(&cates)
		util.CheckErr(err, "查询分类数据错误")
		c.Data["cates"] = cates
		c.TplName = "admin/article-add.html"
	}
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
		util.CheckErr(err, "查询数据错误")
		//赋值更新操作
		arti.Aname = articleName
		arti.Acontent = articleContent
		arti.Aimg = "./static/img/" + fileName
		_, err = o.Update(&arti, "Aname", "Acontent", "Aimg")
		util.CheckErr(err, "更新数据错误")
		c.Redirect("/article/list", 302)

	}
	id, err := c.GetInt("id")
	util.CheckErr(err, "找不到指定的文章id")
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	util.CheckErr(err, "查询失败")
	c.Data["article"] = arti
	c.TplName = "admin/article-update.html"
}

//文章删除
func (c *ArticleController) Delete() {
	id, err := c.GetInt("id")
	util.CheckErr(err, "文章ID获取失败")
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	util.CheckErr(err, "查询错误")
	_, err = o.Delete(&arti)
	util.CheckErr(err, "删除失败")
	c.Redirect("/article/list", 302)
}
