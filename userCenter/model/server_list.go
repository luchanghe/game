package model

type ServerList struct {
	Id     int
	SId    int
	Name   string
	Host   string
	Port   string
	Status int
}

func (ServerList) TableName() string {
	return "server_list"
}
