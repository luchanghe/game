package userManage

import (
	"fmt"
	"game/model"
	"game/tool"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"strconv"
	"strings"
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
			u = model.NewUser()
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
		for _, diff := range r.diffs {
			fmt.Println(diff)
		}
		*watcherStruct.originalUser = *watcherStruct.currentUser
	}
}

type diffReporter struct {
	path  cmp.Path
	diffs []*ChangeCommand
}

func (r *diffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

type ChangeCommand struct {
	Object       string
	Operate      string
	OperateValue string
}

func (r *diffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {

		_, nv := r.path.Last().Values()
		a := strings.Split(r.path.GoString(), ".")
		a = a[1:]
		a[0] = strings.TrimRight(a[0], "}")
		var names []string
		var operate string
		var operateValue string
		for _, f := range a {
			if f[len(f)-1] != ']' {
				names = append(names, f)
				operate = "@s"
				operateValue = getRefValue(nv)
			} else {
				//说明是切片
				names = append(names, f)
				if nv.IsValid() {
					operate = "@s"
					operateValue = getRefValue(nv)
				} else {
					switch r.path.Last().Type().Kind() {
					case reflect.Map:
						operate = "@d"
					case reflect.Slice:
						operate = "@dr"
					default:
						panic("type err")
					}
				}
			}
		}
		r.diffs = append(r.diffs, &ChangeCommand{
			Object:       strings.Join(names, "."),
			Operate:      operate,
			OperateValue: operateValue,
		})
	}
}

func getRefValue(nv reflect.Value) string {
	switch nv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(nv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(nv.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(nv.Float(), 'f', -1, 64)
	case reflect.String:
		return "\"" + nv.String() + "\""
	case reflect.Ptr:
		return getRefValue(nv.Elem())
	case reflect.Struct:
		var buf = strings.Builder{}
		for i := 0; i < nv.NumField(); i++ {
			field := nv.Type().Field(i)
			value := getRefValue(nv.Field(i))
			buf.WriteString(fmt.Sprintf(`{"key":"%s","val":"%s"},`, field.Name, value))
		}
		return strings.TrimRight(buf.String(), ",")
	default:
		return ""
	}
}

func (r *diffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}
