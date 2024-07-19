package userManage

import (
	"bytes"
	"fmt"
	"game/model"
	"game/tool"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"strconv"
	"sync"
)

const ActionUser = "Cyi__ActionUser__"
const UpdateUsers = "Cyi__UpdateUsers__"

var (
	userManageInstance *UserManageStruct
	userManageOnce     sync.Once
)

type UserManageStruct struct {
	users map[int]*model.User
	sync.RWMutex
}

func UserManage() *UserManageStruct {
	userManageOnce.Do(func() {
		userManageInstance = &UserManageStruct{
			users: map[int]*model.User{},
		}
	})
	return userManageInstance
}

func GetUserFromAction(c *gin.Context) *model.User {
	u, _ := c.Get(ActionUser)
	return u.(*model.User)
}

func GetUser(c *gin.Context, userId int) *model.User {
	updateUserMap, ok := c.Get(UpdateUsers)
	if !ok {
		//上下文中不存在结构时创建结构
		updateUserMap = &UpdateUsersStruct{userWatcher: make(map[int]*UpdateUserWatcherStruct)}
		c.Set(UpdateUsers, updateUserMap)
	} else {
		//上下文中存在结构时检查是否有保存的用户对象，如果有则返回
		watcher, ok := updateUserMap.(*UpdateUsersStruct).userWatcher[userId]
		if ok {
			return watcher.currentUser
		}
	}
	//当上下文中没有这个用户对象时去对象缓存中读取并加入上下文
	u := getUserFromGlobalCache(userId)
	uCp := tool.DeepCopy(u)
	updateUserMap.(*UpdateUsersStruct).userWatcher[userId] = &UpdateUserWatcherStruct{currentUser: uCp.(*model.User), originalUser: u}
	return updateUserMap.(*UpdateUsersStruct).userWatcher[userId].currentUser
}

func getUserFromGlobalCache(userId int) *model.User {
	m := UserManage()
	m.RLock()
	u, ok := m.users[userId]
	m.RUnlock()
	if !ok {
		m.Lock()
		defer m.Unlock()
		u, ok = m.users[userId]
		if !ok {
			u = model.NewUser(userId, strconv.Itoa(userId)+"的name", &model.Hero{
				HeroName: "英雄名字",
				HeroId:   userId,
			}, make(map[int]*model.Prop), []int{})
			m.users[userId] = u
		}
	}
	return u
}

type UpdateUsersStruct struct {
	userWatcher map[int]*UpdateUserWatcherStruct
}

type UpdateUserWatcherStruct struct {
	originalUser *model.User
	currentUser  *model.User
}

func GetUserChange(c *gin.Context) {
	updateUserMap, ok := c.Get(UpdateUsers)
	if !ok {
		return
	}
	for _, watcherStruct := range updateUserMap.(*UpdateUsersStruct).userWatcher {
		if cmp.Equal(watcherStruct.originalUser, watcherStruct.currentUser) {
			continue
		}
		var r diffReporter
		opts := cmp.Options{
			cmp.Reporter(&r),
		}
		cmp.Diff(watcherStruct.originalUser, watcherStruct.currentUser, opts)
		//for _, diff := range r.diffs {
		//	//fmt.Println(diff)
		//}
		*watcherStruct.originalUser = *watcherStruct.currentUser
	}
}

type diffReporter struct {
	path  cmp.Path
	diffs []string
}

func (r *diffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *diffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		// 获取比较的对象的类型
		var pathStr bytes.Buffer
		for _, step := range r.path {
			switch s := step.(type) {
			case cmp.StructField:
				if pathStr.Len() > 0 {
					pathStr.WriteString(".")
				}
				pathStr.WriteString(fmt.Sprintf("\"%s\":{\"@s\":}", s.Name()))
			case cmp.SliceIndex:
				if s.Key() == -1 {
					//新增了切片
					_, nv := s.Values()
					pathStr.WriteString(fmt.Sprintf(":{\"@a\":%v}}", nv))
				}
			case cmp.MapIndex:
				pathStr.WriteString(fmt.Sprintf("[\"%v\"]", s.Key()))
			case cmp.TypeAssertion:
				pathStr.WriteString("." + s.String())
			}
		}
		fmt.Println(pathStr.String())
		//_, nv := s.Values()
		//fmt.Println(pathStr.String())
		//i := strings.Index(goStr, "}")
		//d := goStr[i+2:]
		//i = strings.Index(d, ".")
		//if i != -1 {
		//	c := d[i+2]
		//	fmt.Println(c)
		//} else {
		//	i = strings.Index(d, "[")
		//	ov, nv := r.path.Last().Values()
		//	fmt.Printf("修改字段:%s %v => %v", d, ov, nv)
		//}
	}
}

func (r *diffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}
