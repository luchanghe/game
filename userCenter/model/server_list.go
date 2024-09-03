package model

type ServerList struct {
	Id   int
	Name string
	Host string
	Port string
}

func (ServerList) TableName() string {
	return "server_list"
}
