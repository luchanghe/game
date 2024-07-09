package manage

import "sync"

var userLock sync.Map

func isLock(userId int) bool {
	_, ok := userLock.Load(userId)
	return ok
}

func Lock(userId int) bool {
	_, ok := userLock.LoadOrStore(userId, struct{}{})
	return ok
}

func Unlock(userId int) {
	userLock.Delete(userId)
}
