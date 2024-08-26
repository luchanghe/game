package manage

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"server/model"
	"server/tool"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type UserManage struct {
	users      map[int64]*model.User
	nextUserId int64
	sync.RWMutex
}

var userManageOnce sync.Once
var userManageCache *UserManage

func GetUserManage() *UserManage {
	userManageOnce.Do(func() {
		nextUserId := int64(10000000)
		collection := GetMongoManage().GetUserDb().Collection("users")
		opts := options.FindOne().SetSort(bson.D{{"_id", -1}})
		var result bson.M
		filter := bson.M{}
		err := collection.FindOne(context.TODO(), filter, opts).Decode(&result)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				nextUserId = 10000000
			} else {
				log.Fatal(err)
			}
		}
		nextUserId, ok := result["_id"].(int64)
		if !ok {
			log.Fatal("获取最大用户ID时出现异常")
		}
		userManageCache = &UserManage{
			users:      map[int64]*model.User{},
			nextUserId: nextUserId,
		}
	})
	return userManageCache

}

func GetNextUserId() int64 {

	return atomic.AddInt64(&GetUserManage().nextUserId, 1)
}

func GetUserFromAction(c *gin.Context) *model.User {
	u, _ := c.Get(ActionUser)
	return u.(*model.User)
}

func GetUser(c *gin.Context, userId int64) *model.User {
	updateUserMap, ok := c.Get(UpdateUsers)
	if !ok {
		//上下文中不存在结构时创建结构
		updateUserMap = &UpdateUsersStruct{userWatcher: make(map[int64]*UpdateUserWatcherStruct)}
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

func getUserFromGlobalCache(userId int64) *model.User {
	m := GetUserManage()
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
	userWatcher map[int64]*UpdateUserWatcherStruct
}

type UpdateUserWatcherStruct struct {
	originalUser *model.User
	currentUser  *model.User
}

func GetUserChange(c *gin.Context) map[int64][]*ChangeCommand {
	updateUserMap, ok := c.Get(UpdateUsers)
	changeMap := make(map[int64][]*ChangeCommand)
	if !ok {
		return changeMap
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
		changeMap[watcherStruct.originalUser.Id] = r.diffs
		*watcherStruct.originalUser = *watcherStruct.currentUser
	}
	return changeMap
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
				if i := strings.Index(f, "?->"); i != -1 {
					f = f[:i] + f[i+3:]
				}
				if i := strings.Index(f, "->?"); i != -1 {
					f = f[:i] + f[i+3:]
				}
				if i := strings.Index(f, "["); i != -1 {
					f = f[:i] + "." + f[i+1:len(f)-1]
				}
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
		return nv.String()
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
