package models

type Category struct {
	Id int
	TypeName string `orm:"size(20)"`
}

func (m *Category) TableName() string {
	return TableName("category")
}