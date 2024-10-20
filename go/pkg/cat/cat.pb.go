// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: cat/cat.proto

package cat

import (
	core "github.com/mojo-lang/core/go/pkg/mojo/core"
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

type Cat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NickName      string          `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nickName,omitempty"`
	Sex           string          `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	BrithDay      *core.Timestamp `protobuf:"bytes,5,opt,name=brith_day,json=brithDay,proto3" json:"brithDay,omitempty"`
	Location      string          `protobuf:"bytes,20,opt,name=location,proto3" json:"location,omitempty"`
	Area          string          `protobuf:"bytes,21,opt,name=area,proto3" json:"area,omitempty"`
	LastFoundTime *core.Timestamp `protobuf:"bytes,25,opt,name=last_found_time,json=lastFoundTime,proto3" json:"lastFoundTime,omitempty"`
	Status        string          `protobuf:"bytes,50,opt,name=status,proto3" json:"status,omitempty"`
	State         string          `protobuf:"bytes,55,opt,name=state,proto3" json:"state,omitempty"`
	CreateTime    *core.Timestamp `protobuf:"bytes,100,opt,name=create_time,json=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime    *core.Timestamp `protobuf:"bytes,101,opt,name=update_time,json=updateTime,proto3" json:"updateTime,omitempty"`
	DeleteTime    *core.Timestamp `protobuf:"bytes,102,opt,name=delete_time,json=deleteTime,proto3" json:"deleteTime,omitempty"`
}

func (x *Cat) Reset() {
	*x = Cat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cat_cat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cat) ProtoMessage() {}

func (x *Cat) ProtoReflect() protoreflect.Message {
	mi := &file_cat_cat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cat.ProtoReflect.Descriptor instead.
func (*Cat) Descriptor() ([]byte, []int) {
	return file_cat_cat_proto_rawDescGZIP(), []int{0}
}

func (x *Cat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cat) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *Cat) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *Cat) GetBrithDay() *core.Timestamp {
	if x != nil {
		return x.BrithDay
	}
	return nil
}

func (x *Cat) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Cat) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *Cat) GetLastFoundTime() *core.Timestamp {
	if x != nil {
		return x.LastFoundTime
	}
	return nil
}

func (x *Cat) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Cat) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Cat) GetCreateTime() *core.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Cat) GetUpdateTime() *core.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Cat) GetDeleteTime() *core.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

var File_cat_cat_proto protoreflect.FileDescriptor

var file_cat_cat_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x74, 0x2f, 0x63, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x63, 0x61, 0x74, 0x1a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x03, 0x0a, 0x03, 0x43,
	0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x31, 0x0a, 0x09, 0x62, 0x72, 0x69, 0x74, 0x68, 0x5f, 0x64,
	0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x62, 0x72, 0x69, 0x74, 0x68, 0x44, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x3c, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x37, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x44, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x6c, 0x69, 0x61, 0x6e, 0x6b, 0x75, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x42, 0x08, 0x43, 0x61,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x61, 0x6e, 0x6b, 0x75, 0x69, 0x2f, 0x74, 0x6e, 0x72,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x74, 0x3b, 0x63, 0x61, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cat_cat_proto_rawDescOnce sync.Once
	file_cat_cat_proto_rawDescData = file_cat_cat_proto_rawDesc
)

func file_cat_cat_proto_rawDescGZIP() []byte {
	file_cat_cat_proto_rawDescOnce.Do(func() {
		file_cat_cat_proto_rawDescData = protoimpl.X.CompressGZIP(file_cat_cat_proto_rawDescData)
	})
	return file_cat_cat_proto_rawDescData
}

var file_cat_cat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cat_cat_proto_goTypes = []interface{}{
	(*Cat)(nil),            // 0: cat.Cat
	(*core.Timestamp)(nil), // 1: mojo.core.Timestamp
}
var file_cat_cat_proto_depIdxs = []int32{
	1, // 0: cat.Cat.brith_day:type_name -> mojo.core.Timestamp
	1, // 1: cat.Cat.last_found_time:type_name -> mojo.core.Timestamp
	1, // 2: cat.Cat.create_time:type_name -> mojo.core.Timestamp
	1, // 3: cat.Cat.update_time:type_name -> mojo.core.Timestamp
	1, // 4: cat.Cat.delete_time:type_name -> mojo.core.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cat_cat_proto_init() }
func file_cat_cat_proto_init() {
	if File_cat_cat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cat_cat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cat); i {
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
			RawDescriptor: file_cat_cat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cat_cat_proto_goTypes,
		DependencyIndexes: file_cat_cat_proto_depIdxs,
		MessageInfos:      file_cat_cat_proto_msgTypes,
	}.Build()
	File_cat_cat_proto = out.File
	file_cat_cat_proto_rawDesc = nil
	file_cat_cat_proto_goTypes = nil
	file_cat_cat_proto_depIdxs = nil
}