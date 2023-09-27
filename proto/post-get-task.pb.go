// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/post-get-task.proto

package proto

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

type PostGetTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskName string `protobuf:"bytes,1,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
}

func (x *PostGetTask) Reset() {
	*x = PostGetTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_get_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostGetTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostGetTask) ProtoMessage() {}

func (x *PostGetTask) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_get_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostGetTask.ProtoReflect.Descriptor instead.
func (*PostGetTask) Descriptor() ([]byte, []int) {
	return file_proto_post_get_task_proto_rawDescGZIP(), []int{0}
}

func (x *PostGetTask) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

var File_proto_post_get_task_proto protoreflect.FileDescriptor

var file_proto_post_get_task_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2d, 0x67, 0x65, 0x74,
	0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x2a, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x19, 0x5a,
	0x17, 0x74, 0x6f, 0x64, 0x6f, 0x77, 0x68, 0x65, 0x65, 0x6c, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_post_get_task_proto_rawDescOnce sync.Once
	file_proto_post_get_task_proto_rawDescData = file_proto_post_get_task_proto_rawDesc
)

func file_proto_post_get_task_proto_rawDescGZIP() []byte {
	file_proto_post_get_task_proto_rawDescOnce.Do(func() {
		file_proto_post_get_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_post_get_task_proto_rawDescData)
	})
	return file_proto_post_get_task_proto_rawDescData
}

var file_proto_post_get_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_post_get_task_proto_goTypes = []interface{}{
	(*PostGetTask)(nil), // 0: main.PostGetTask
}
var file_proto_post_get_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_post_get_task_proto_init() }
func file_proto_post_get_task_proto_init() {
	if File_proto_post_get_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_post_get_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostGetTask); i {
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
			RawDescriptor: file_proto_post_get_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_post_get_task_proto_goTypes,
		DependencyIndexes: file_proto_post_get_task_proto_depIdxs,
		MessageInfos:      file_proto_post_get_task_proto_msgTypes,
	}.Build()
	File_proto_post_get_task_proto = out.File
	file_proto_post_get_task_proto_rawDesc = nil
	file_proto_post_get_task_proto_goTypes = nil
	file_proto_post_get_task_proto_depIdxs = nil
}
