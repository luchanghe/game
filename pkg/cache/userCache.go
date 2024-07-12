package cache

import (
	"game/been/redis/user"
	"sync"
)

var userCache = sync.Map{}

func GetUser(userId int) *user.User {
	u, ok := userCache.Load(userId)
	if !ok {
		u = &user.User{Id: userId, Name: "假装是数据库读取的名字"}
		userCache.Store(userId, &user.User{Id: userId, Name: "从数据库读取的名字"})
	}
	return u.(*user.User)
}
