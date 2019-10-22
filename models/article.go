package models

import "time"

type Article struct {
	Id       int       `orm:"pk;auto"`  //主机和自增
	Aname    string    `orm:"size(50)"` //长度50
	Atime    time.Time `orm:"auto_now"`
	Acount   int       `orm:default(0);null` //默认值0 允许为空
	Acontent string
	Aimg     string
}

func (m *Article) TableName() string {
	return TableName("article")
}
