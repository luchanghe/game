package model

import (
	"server/tool"
)

type User struct {
	Id        int64         `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name,omitempty" json:"name"`
	Hero      *Hero         `bson:"hero,omitempty" json:"hero"`
	Props     map[int]*Prop `bson:"props,omitempty" json:"props"`
	NormalInt []int         `bson:"normalInt,omitempty" json:"normalInt"`
}

func NewUser() *User {
	m := &User{}
	tool.InitStruct(m)
	return m
}
