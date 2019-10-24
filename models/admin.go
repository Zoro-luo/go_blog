package models

type Admin struct {
	Id       int
	UserName string `orm:"unique"`
	Password string
}
