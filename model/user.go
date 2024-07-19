package model

type User struct {
	Id        int
	Name      string
	Hero      *Hero
	Props     map[int]*Prop
	NormalInt []int
}

func NewUser(id int, name string, hero *Hero, props map[int]*Prop, normalInt []int) *User {
	return &User{Id: id, Name: name, Hero: hero, Props: props, NormalInt: normalInt}
}
