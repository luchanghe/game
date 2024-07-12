package user

import "game/been/redis"

type User struct {
	Id      int
	Name    string
	Watcher *redis.Watcher
}

func NewUser(id int, name string) *User {
	u := &User{
		Id:   id,
		Name: name,
	}
	u.Watcher = redis.NewWatcher(u)
	return u
}
