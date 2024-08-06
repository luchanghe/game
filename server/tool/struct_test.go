package tool

import (
	"fmt"
	"testing"
)

type MyStruct struct {
	MyMap map[string]int
}

func TestInitStruct(t *testing.T) {
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

	u := user{}
	InitStruct(&u)
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
	fmt.Println(u)
}
