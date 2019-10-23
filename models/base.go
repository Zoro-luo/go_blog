package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //mysql 驱动
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	//设置数据库基本信息
	orm.RegisterDataBase(`default`, "mysql", dsn)
	//指定model映射表
	orm.RegisterModel(new(User), new(Admin), new(Article),new(Category))
	//生成表
	//name=>default [别名] force=>false[是否强制更新表结构] verbose=>true[是否可见创建sql过程]
	orm.RunSyncdb("default", false, true)
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
