package main

import (
	"github.com/astaxie/beego"
	_ "liteblog/routers"
	"strconv"
	"strings"
	"liteblog/models"
)

func init()  {
	//models包调用生成映射表
	models.Init()
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.AddFuncMap("ShowPrePage",HandlePrepage)		//视图函数 上一页
	beego.AddFuncMap("ShowNextPage",HandleNextpage)		//视图函数 上一页
	initTemlate()
	beego.Run()
}

//比较两个url字符串 相同则返回true [调用在:view/home/common/header.html]
func initTemlate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		urlX := strings.Trim(x, "/")
		urlY := strings.Trim(y, "/")
		return strings.Compare(urlX, urlY) == 0
	})
}

//上一页
func HandlePrepage(data int) string {
	pageIndex := data - 1
	return strconv.Itoa(pageIndex)
}
//下一页
func HandleNextpage(data int) string {
	pageIndex := data + 1
	return strconv.Itoa(pageIndex)
}
