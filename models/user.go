package models

type User struct {
	Id   int
	Name string `orm:"unique"`
	Pwd  string
}

func (m *User) TableName() string {
	return TableName("user")
}
