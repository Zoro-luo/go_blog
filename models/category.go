package models

type Category struct {
	Id int
	TypeName string `orm:"size(20)"`
}
