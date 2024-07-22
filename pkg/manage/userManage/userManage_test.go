package userManage

import (
	"fmt"
	"game/model"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGet(t *testing.T) {
	c := &gin.Context{}
	u := GetUser(c, 100001)
	u2 := GetUser(c, 100002)
	fmt.Println(u, u2)
}

func TestUserCp(t *testing.T) {
	c := &gin.Context{}
	u := GetUser(c, 100001)
	u2 := GetUser(c, 100001)
	updateUserMap, _ := c.Get(UpdateUsers)
	u.Name = "firstName"
	u.Hero.HeroId = 1000
	u.Props[10] = &model.Prop{PropId: 10, PropNum: 10}
	fmt.Println("对象池", UserManage().users[100001], "上下文中的原始数据", updateUserMap.(*UpdateUsersStruct).userWatcher[100001].originalUser, "上下文中的当前数据", updateUserMap.(*UpdateUsersStruct).userWatcher[100001].currentUser, "二次获得", u2)
}
