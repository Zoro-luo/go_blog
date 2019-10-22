package main

import (
	"github.com/astaxie/beego"
	_ "liteblog/routers"
	"strings"
	"liteblog/models"
)

func init()  {
	//models包调用生成映射表
	models.Init()
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
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

