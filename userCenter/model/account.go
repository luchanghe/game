package model

type User struct {
	Id       int
	Account  string
	Password string
}

func (User) TableName() string {
	return "users"
}
