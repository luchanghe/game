package user

import "game/pkg/change"

type User struct {
	Id      int
	Name    string
	Seats   []HeroSeat
	Watcher *change.Watcher
}

func NewUser(id int, name string) *User {
	u := &User{
		Id:   id,
		Name: name,
	}
	u.Watcher = change.NewWatcher(u)
	return u
}
