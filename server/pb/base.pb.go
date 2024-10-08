// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: base.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId   int32       `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	HeroName string      `protobuf:"bytes,2,opt,name=HeroName,proto3" json:"HeroName,omitempty"`
	HeroAttr []*HeroAttr `protobuf:"bytes,3,rep,name=HeroAttr,proto3" json:"HeroAttr,omitempty"`
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *Hero) GetHeroId() int32 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *Hero) GetHeroName() string {
	if x != nil {
		return x.HeroName
	}
	return ""
}

func (x *Hero) GetHeroAttr() []*HeroAttr {
	if x != nil {
		return x.HeroAttr
	}
	return nil
}

type HeroAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrId int32 `protobuf:"varint,1,opt,name=AttrId,proto3" json:"AttrId,omitempty"`
	Value  int32 `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *HeroAttr) Reset() {
	*x = HeroAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeroAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeroAttr) ProtoMessage() {}

func (x *HeroAttr) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeroAttr.ProtoReflect.Descriptor instead.
func (*HeroAttr) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *HeroAttr) GetAttrId() int32 {
	if x != nil {
		return x.AttrId
	}
	return 0
}

func (x *HeroAttr) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Prop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropId  int32 `protobuf:"varint,1,opt,name=PropId,proto3" json:"PropId,omitempty"`
	PropNum int32 `protobuf:"varint,2,opt,name=PropNum,proto3" json:"PropNum,omitempty"`
}

func (x *Prop) Reset() {
	*x = Prop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prop) ProtoMessage() {}

func (x *Prop) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prop.ProtoReflect.Descriptor instead.
func (*Prop) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *Prop) GetPropId() int32 {
	if x != nil {
		return x.PropId
	}
	return 0
}

func (x *Prop) GetPropNum() int32 {
	if x != nil {
		return x.PropNum
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64           `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string          `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Hero      *Hero           `protobuf:"bytes,3,opt,name=Hero,proto3" json:"Hero,omitempty"`
	Props     map[int32]*Prop `protobuf:"bytes,4,rep,name=Props,proto3" json:"Props,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NormalInt []int32         `protobuf:"varint,5,rep,packed,name=NormalInt,proto3" json:"NormalInt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetHero() *Hero {
	if x != nil {
		return x.Hero
	}
	return nil
}

func (x *User) GetProps() map[int32]*Prop {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *User) GetNormalInt() []int32 {
	if x != nil {
		return x.NormalInt
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x64, 0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x48, 0x65, 0x72, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x48, 0x65, 0x72, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x08,
	0x48, 0x65, 0x72, 0x6f, 0x41, 0x74, 0x74, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x41, 0x74, 0x74, 0x72, 0x52, 0x08, 0x48, 0x65,
	0x72, 0x6f, 0x41, 0x74, 0x74, 0x72, 0x22, 0x38, 0x0a, 0x08, 0x48, 0x65, 0x72, 0x6f, 0x41, 0x74,
	0x74, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x74, 0x74, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x41, 0x74, 0x74, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x38, 0x0a, 0x04, 0x50, 0x72, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x70,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x70, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x70, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x70, 0x4e, 0x75, 0x6d, 0x22, 0xd5, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52,
	0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x09, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x1a, 0x42,
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_base_proto_goTypes = []any{
	(*Hero)(nil),     // 0: pb.Hero
	(*HeroAttr)(nil), // 1: pb.HeroAttr
	(*Prop)(nil),     // 2: pb.Prop
	(*User)(nil),     // 3: pb.User
	nil,              // 4: pb.User.PropsEntry
}
var file_base_proto_depIdxs = []int32{
	1, // 0: pb.Hero.HeroAttr:type_name -> pb.HeroAttr
	0, // 1: pb.User.Hero:type_name -> pb.Hero
	4, // 2: pb.User.Props:type_name -> pb.User.PropsEntry
	2, // 3: pb.User.PropsEntry.value:type_name -> pb.Prop
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Hero); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_base_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HeroAttr); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_base_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Prop); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_base_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
