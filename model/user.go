package model

import "game/tool"

type User struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	Hero      *Hero         `json:"hero"`
	Props     map[int]*Prop `json:"props"`
	NormalInt []int         `json:"normalInt"`
}

func NewUser() *User {
	m := &User{}
	tool.InitStruct(m)
	return m
}
