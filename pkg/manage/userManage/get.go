package userManage

import (
	"game/model/user"
	"github.com/gin-gonic/gin"
	"sync"
)

var users sync.Map
var usersMapLock sync.Mutex

func GetUserFormUid(userId int) (*user.User, error) {
	if !usersMapLock.TryLock() {
		panic("加锁异常，后续要增加等待或者其他措施")
	}
	u, ok := users.Load(userId)
	if !ok {
		//这里应该去读数据，目前先伪造一下
		u = user.NewUser(100000, "Cyi", nil)
		users.Store(userId, u)
	}
	usersMapLock.Unlock()
	return u.(*user.User), nil
}

func GetUserFromAction(c *gin.Context) *user.User {
	u, _ := c.Get("user")
	return u.(*user.User)
}
