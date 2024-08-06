package model

import (
	"server/tool"
)

type User struct {
	Id        int64         `bson:"_id,omitempty"`
	Name      string        `bson:"name,omitempty"`
	Hero      *Hero         `bson:"hero,omitempty"`
	Props     map[int]*Prop `bson:"props,omitempty"`
	NormalInt []int         `bson:"normalInt,omitempty"`
}

func NewUser() *User {
	m := &User{}
	tool.InitStruct(m)
	return m
}
