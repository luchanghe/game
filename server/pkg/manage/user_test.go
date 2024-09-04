package manage

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"server/model"
	"server/pb"
	"server/pkg/sysConst"
	"server/tool"
	"testing"
)

func TestGet(t *testing.T) {
	c := &gin.Context{}
	u, _ := GetUser(c, 100001)
	u2, _ := GetUser(c, 100002)
	fmt.Println(u, u2)
}

func TestUserCp(t *testing.T) {
	c := &gin.Context{}
	u, _ := GetUser(c, 100001)
	u2, _ := GetUser(c, 100001)
	updateUserMap, _ := c.Get(sysDefined.UpdateUsers)
	u.Name = "firstName"
	u.Hero.HeroId = 1000
	u.Props[10] = &model.Prop{PropId: 10, PropNum: 10}
	fmt.Println("对象池", GetUserManage().users[100001], "上下文中的原始数据", updateUserMap.(*UpdateUsersStruct).userWatcher[100001].originalUser, "上下文中的当前数据", updateUserMap.(*UpdateUsersStruct).userWatcher[100001].currentUser, "二次获得", u2)
}

type userTestStruct struct {
	Int    int
	String string
	Object *userTestStruct
}
type user struct {
	Int_Int_Map        map[int]int
	Int_String_Map     map[int]string
	Int_Struct_Ptr_Map map[int]*userTestStruct
	Int_Struct_Map     map[int]userTestStruct

	String_Int_Map        map[string]int
	String_String_Map     map[string]string
	String_Struct_Ptr_Map map[string]*userTestStruct
	String_Struct_Map     map[string]userTestStruct

	Int_Slice        []int
	String_Slice     []string
	Struct_Ptr_Slice []*userTestStruct
	Struct_Slice     []userTestStruct
}

func TestUserChangeMapAdd(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Int_Map:           make(map[int]int),
		Int_String_Map:        make(map[int]string),
		Int_Struct_Ptr_Map:    make(map[int]*userTestStruct),
		Int_Struct_Map:        make(map[int]userTestStruct),
		String_Int_Map:        make(map[string]int),
		String_String_Map:     make(map[string]string),
		String_Struct_Ptr_Map: make(map[string]*userTestStruct),
		String_Struct_Map:     make(map[string]userTestStruct),
	}
	u2 := tool.DeepCopy(u)
	u2.(*user).Int_Int_Map[1] = 1
	u2.(*user).Int_String_Map[1] = "1"
	u2.(*user).Int_Struct_Ptr_Map[1] = &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u2.(*user).Int_Struct_Map[1] = userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u2.(*user).String_Int_Map["1"] = 1
	u2.(*user).String_String_Map["1"] = "1"
	u2.(*user).String_Struct_Ptr_Map["1"] = &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	cmp.Diff(u, u2, opts)
}

func TestUserChangeMapUpdate(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Int_Map:           make(map[int]int),
		Int_String_Map:        make(map[int]string),
		Int_Struct_Ptr_Map:    make(map[int]*userTestStruct),
		Int_Struct_Map:        make(map[int]userTestStruct),
		String_Int_Map:        make(map[string]int),
		String_String_Map:     make(map[string]string),
		String_Struct_Ptr_Map: make(map[string]*userTestStruct),
		String_Struct_Map:     make(map[string]userTestStruct)}
	u.Int_Int_Map[1] = 1
	u.Int_String_Map[1] = "1"
	u.Int_Struct_Ptr_Map[1] = &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u.Int_Struct_Map[1] = userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u.String_Int_Map["1"] = 1
	u.String_String_Map["1"] = "1"
	u.String_Struct_Ptr_Map["1"] = &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u2 := tool.DeepCopy(u)
	u2.(*user).Int_Int_Map[1] = 2
	u2.(*user).Int_String_Map[1] = "2"
	u2.(*user).Int_Struct_Ptr_Map[1] = &userTestStruct{
		Int:    1,
		String: "测试2",
		Object: &userTestStruct{
			Int:    2,
			String: "测试2层",
			Object: nil,
		},
	}
	u2.(*user).String_Int_Map["1"] = 2
	u2.(*user).String_String_Map["1"] = "2"
	u2.(*user).String_Struct_Ptr_Map["1"] = &userTestStruct{
		Int:    1,
		String: "测试2",
		Object: &userTestStruct{
			Int:    2,
			String: "测试2层",
			Object: nil,
		},
	}
	u2.(*user).String_Struct_Map["1"] = userTestStruct{
		Int:    1,
		String: "测试2",
		Object: &userTestStruct{
			Int:    2,
			String: "测试2层",
			Object: nil,
		},
	}
	cmp.Diff(u, u2, opts)
}

func TestUserChangeMapDelete(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Int_Map:           make(map[int]int),
		Int_String_Map:        make(map[int]string),
		Int_Struct_Ptr_Map:    make(map[int]*userTestStruct),
		Int_Struct_Map:        make(map[int]userTestStruct),
		String_Int_Map:        make(map[string]int),
		String_String_Map:     make(map[string]string),
		String_Struct_Ptr_Map: make(map[string]*userTestStruct),
		String_Struct_Map:     make(map[string]userTestStruct)}
	u.Int_Int_Map[1] = 1
	u.Int_String_Map[1] = "1"
	u.Int_Struct_Ptr_Map[1] = &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u.Int_Struct_Map[1] = userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u.String_Int_Map["1"] = 1
	u.String_String_Map["1"] = "1"
	u.String_Struct_Ptr_Map["1"] = &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u.String_Struct_Map["1"] = userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u2 := tool.DeepCopy(u)
	delete(u2.(*user).Int_Int_Map, 1)
	delete(u2.(*user).Int_String_Map, 1)
	delete(u2.(*user).Int_Struct_Ptr_Map, 1)
	delete(u2.(*user).String_Int_Map, "1")
	delete(u2.(*user).String_String_Map, "1")
	delete(u2.(*user).String_Struct_Ptr_Map, "1")
	cmp.Diff(u, u2, opts)
}

func TestUserChangeSliceAdd(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Slice:        make([]int, 0),
		String_Slice:     make([]string, 0),
		Struct_Slice:     make([]userTestStruct, 0),
		Struct_Ptr_Slice: make([]*userTestStruct, 0)}
	u2 := tool.DeepCopy(u)
	u2.(*user).Int_Slice = append(u2.(*user).Int_Slice, 1)
	u2.(*user).String_Slice = append(u2.(*user).String_Slice, "1")
	u2.(*user).Struct_Slice = append(u2.(*user).Struct_Slice, userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2.(*user).Struct_Ptr_Slice = append(u2.(*user).Struct_Ptr_Slice, &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	cmp.Diff(u, u2, opts)
}

func TestUserChangeSliceUpdate(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Slice:        make([]int, 0),
		String_Slice:     make([]string, 0),
		Struct_Slice:     make([]userTestStruct, 0),
		Struct_Ptr_Slice: make([]*userTestStruct, 0)}
	u.Int_Slice = append(u.Int_Slice, 1)
	u.String_Slice = append(u.String_Slice, "1")
	u.Struct_Slice = append(u.Struct_Slice, userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u.Struct_Ptr_Slice = append(u.Struct_Ptr_Slice, &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2 := tool.DeepCopy(u)
	u2.(*user).Int_Slice[0] = 2
	u2.(*user).String_Slice[0] = "2"
	u2.(*user).Struct_Slice[0] = userTestStruct{
		Int:    6,
		String: "测试2",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	u2.(*user).Struct_Ptr_Slice[0] = &userTestStruct{
		Int:    4,
		String: "测试2",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}
	cmp.Diff(u, u2, opts)
}

func TestUserChangeSliceDelete(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Slice:        make([]int, 0),
		String_Slice:     make([]string, 0),
		Struct_Slice:     make([]userTestStruct, 0),
		Struct_Ptr_Slice: make([]*userTestStruct, 0)}
	u.Int_Slice = append(u.Int_Slice, 1)
	u.String_Slice = append(u.String_Slice, "1")
	u.Struct_Slice = append(u.Struct_Slice, userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u.Struct_Ptr_Slice = append(u.Struct_Ptr_Slice, &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2 := tool.DeepCopy(u)
	u2.(*user).Int_Slice = u2.(*user).Int_Slice[1:]
	u2.(*user).String_Slice = u2.(*user).String_Slice[1:]
	u2.(*user).Struct_Slice = u2.(*user).Struct_Slice[1:]
	u2.(*user).Struct_Ptr_Slice = u2.(*user).Struct_Ptr_Slice[1:]
	cmp.Diff(u, u2, opts)
}

func TestUserChangeSliceDeleteAndAdd(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Slice:        make([]int, 0),
		String_Slice:     make([]string, 0),
		Struct_Slice:     make([]userTestStruct, 0),
		Struct_Ptr_Slice: make([]*userTestStruct, 0)}
	u.Int_Slice = append(u.Int_Slice, 1)
	u.String_Slice = append(u.String_Slice, "1")
	u.Struct_Slice = append(u.Struct_Slice, userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u.Struct_Ptr_Slice = append(u.Struct_Ptr_Slice, &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2 := tool.DeepCopy(u)
	u2.(*user).Int_Slice = u2.(*user).Int_Slice[1:]
	u2.(*user).String_Slice = u2.(*user).String_Slice[1:]
	u2.(*user).Struct_Slice = u2.(*user).Struct_Slice[1:]
	u2.(*user).Struct_Ptr_Slice = u2.(*user).Struct_Ptr_Slice[1:]

	u2.(*user).Int_Slice = append(u2.(*user).Int_Slice, 6)
	u2.(*user).String_Slice = append(u2.(*user).String_Slice, "6")
	u2.(*user).Struct_Slice = append(u2.(*user).Struct_Slice, userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2.(*user).Struct_Ptr_Slice = append(u2.(*user).Struct_Ptr_Slice, &userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}, &userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	cmp.Diff(u, u2, opts)
}

func TestAddToChangeRes(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Slice:        make([]int, 0),
		String_Slice:     make([]string, 0),
		Struct_Slice:     make([]userTestStruct, 0),
		Struct_Ptr_Slice: make([]*userTestStruct, 0)}
	u.Int_Slice = append(u.Int_Slice, 1)
	u.String_Slice = append(u.String_Slice, "1")
	u.Struct_Slice = append(u.Struct_Slice, userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u.Struct_Ptr_Slice = append(u.Struct_Ptr_Slice, &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2 := tool.DeepCopy(u)
	u2.(*user).Int_Slice = u2.(*user).Int_Slice[1:]
	u2.(*user).String_Slice = u2.(*user).String_Slice[1:]
	u2.(*user).Struct_Slice = u2.(*user).Struct_Slice[1:]
	u2.(*user).Struct_Ptr_Slice = u2.(*user).Struct_Ptr_Slice[1:]

	u2.(*user).Int_Slice = append(u2.(*user).Int_Slice, 6)
	u2.(*user).String_Slice = append(u2.(*user).String_Slice, "6")
	u2.(*user).Struct_Slice = append(u2.(*user).Struct_Slice, userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2.(*user).Struct_Ptr_Slice = append(u2.(*user).Struct_Ptr_Slice, &userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}, &userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	cmp.Diff(u, u2, opts)
	var res proto.Message
	res = &pb.DefaultResponse{}
	fieldDescriptor := res.ProtoReflect().Descriptor().Fields().ByName("c")
	changeMessage := &pb.ChangeMessage{
		ChangeCommand: []*pb.ChangeMessage_Command{},
	}
	for _, diff := range r.diffs {
		changeMessage.ChangeCommand = append(changeMessage.ChangeCommand, &pb.ChangeMessage_Command{Object: diff.Object, Operate: diff.Operate, OperateValue: diff.OperateValue})
	}
	changeMessageValue := protoreflect.ValueOf(changeMessage.ProtoReflect())
	res.ProtoReflect().Set(fieldDescriptor, changeMessageValue)
	fmt.Println(res)
}

func TestGetChangeToDb(t *testing.T) {
	var r diffReporter
	opts := cmp.Options{
		cmp.Reporter(&r),
	}
	u := &user{
		Int_Slice:        make([]int, 0),
		String_Slice:     make([]string, 0),
		Struct_Slice:     make([]userTestStruct, 0),
		Struct_Ptr_Slice: make([]*userTestStruct, 0)}
	u.Int_Slice = append(u.Int_Slice, 1)
	u.String_Slice = append(u.String_Slice, "1")
	u.Struct_Slice = append(u.Struct_Slice, userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u.Struct_Ptr_Slice = append(u.Struct_Ptr_Slice, &userTestStruct{
		Int:    1,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2 := tool.DeepCopy(u)
	u2.(*user).Int_Slice = u2.(*user).Int_Slice[1:]
	u2.(*user).String_Slice = u2.(*user).String_Slice[1:]
	u2.(*user).Struct_Slice = u2.(*user).Struct_Slice[1:]
	u2.(*user).Struct_Ptr_Slice = u2.(*user).Struct_Ptr_Slice[1:]

	u2.(*user).Int_Slice = append(u2.(*user).Int_Slice, 6)
	u2.(*user).String_Slice = append(u2.(*user).String_Slice, "6")
	u2.(*user).Struct_Slice = append(u2.(*user).Struct_Slice, userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2.(*user).Struct_Ptr_Slice = append(u2.(*user).Struct_Ptr_Slice, &userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	}, &userTestStruct{
		Int:    3,
		String: "测试1",
		Object: &userTestStruct{
			Int:    0,
			String: "测试2层",
			Object: nil,
		},
	})
	u2.(*user).Struct_Ptr_Slice = append(u2.(*user).Struct_Ptr_Slice[1:])
	cmp.Diff(u, u2, opts)
	setFields := bson.D{}
	unsetFields := bson.D{}
	for _, diff := range r.diffs {
		//diff.Object 是key diff.OperateValue是value
		fmt.Println(diff.Object, diff.Operate, diff.OperateValue)
		if diff.Operate == "@s" {
			//如何加入到$set
			setFields = append(setFields, bson.E{Key: diff.Object, Value: diff.OperateValue})
		} else {
			unsetFields = append(unsetFields, bson.E{Key: diff.Object, Value: ""})
		}
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.TODO())
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
		return
	}
	fmt.Println("Connected to MongoDB!")
	// 获取数据库和集合
	collection := client.Database("testdb").Collection("testcollection")
	// 创建更新操作
	update := bson.D{}
	if len(setFields) > 0 {
		update = append(update, bson.E{Key: "$set", Value: setFields})
	}
	if len(unsetFields) > 0 {
		update = append(update, bson.E{Key: "$unset", Value: unsetFields})
	}
	filter := bson.D{{"_id", 1}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Error updating document:", err)
		return
	}
	fmt.Printf("Matched %d document(s) and updated %d document(s).\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

type User struct {
	ID      uint64 `bson:"_id,omitempty"`
	Name    string `bson:"name,omitempty"`
	Age     int    `bson:"age,omitempty"`
	Email   string `bson:"email,omitempty"`
	Address string `bson:"address,omitempty"`
}
