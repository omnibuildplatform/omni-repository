// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0--rc1
// source: call_repo.proto

package app

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

// The request message containing the user's name.
type RepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	UserId   int32  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *RepRequest) Reset() {
	*x = RepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepRequest) ProtoMessage() {}

func (x *RepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_call_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepRequest.ProtoReflect.Descriptor instead.
func (*RepRequest) Descriptor() ([]byte, []int) {
	return file_call_repo_proto_rawDescGZIP(), []int{0}
}

func (x *RepRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RepRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RepRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// The response message containing the greetings
type RepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RepResponse) Reset() {
	*x = RepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_call_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepResponse) ProtoMessage() {}

func (x *RepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_call_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepResponse.ProtoReflect.Descriptor instead.
func (*RepResponse) Descriptor() ([]byte, []int) {
	return file_call_repo_proto_rawDescGZIP(), []int{1}
}

func (x *RepResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RepResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_call_repo_proto protoreflect.FileDescriptor

var file_call_repo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x70, 0x22, 0x68, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x3f, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x41, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x33, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x0f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x52, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x52, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x4f, 0x6d, 0x6e, 0x69, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_call_repo_proto_rawDescOnce sync.Once
	file_call_repo_proto_rawDescData = file_call_repo_proto_rawDesc
)

func file_call_repo_proto_rawDescGZIP() []byte {
	file_call_repo_proto_rawDescOnce.Do(func() {
		file_call_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_call_repo_proto_rawDescData)
	})
	return file_call_repo_proto_rawDescData
}

var file_call_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_call_repo_proto_goTypes = []interface{}{
	(*RepRequest)(nil),  // 0: app.RepRequest
	(*RepResponse)(nil), // 1: app.RepResponse
}
var file_call_repo_proto_depIdxs = []int32{
	0, // 0: app.RepoServer.CallLoadFrom:input_type -> app.RepRequest
	1, // 1: app.RepoServer.CallLoadFrom:output_type -> app.RepResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_call_repo_proto_init() }
func file_call_repo_proto_init() {
	if File_call_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_call_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepRequest); i {
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
		file_call_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepResponse); i {
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
			RawDescriptor: file_call_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_call_repo_proto_goTypes,
		DependencyIndexes: file_call_repo_proto_depIdxs,
		MessageInfos:      file_call_repo_proto_msgTypes,
	}.Build()
	File_call_repo_proto = out.File
	file_call_repo_proto_rawDesc = nil
	file_call_repo_proto_goTypes = nil
	file_call_repo_proto_depIdxs = nil
}