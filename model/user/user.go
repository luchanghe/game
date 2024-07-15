package user

import "game/model/hero"

type User struct {
	Id   int
	Name string
	Hero []hero.Hero
}

func NewUser(id int, name string, hero []hero.Hero) *User {
	return &User{Id: id, Name: name, Hero: hero}
}
