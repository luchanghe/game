package userManage

import "sync"

var userLock sync.Map

func isLock(userId int64) bool {
	_, ok := userLock.Load(userId)
	return ok
}

func Lock(userId int64) bool {
	_, ok := userLock.LoadOrStore(userId, struct{}{})
	return ok
}

func Unlock(userId int64) {
	userLock.Delete(userId)
}
