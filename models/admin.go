package models

type Admin struct {
	Id       int
	UserName string `orm:"unique"`
	Password string
}

func (m *Admin) TableName() string {
	return TableName("admin")
}
