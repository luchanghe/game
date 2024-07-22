package userManage

import (
	"game/tool"
	"github.com/google/go-cmp/cmp"
	"testing"
)

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
