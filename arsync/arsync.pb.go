// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: arsync/arsync.proto

package arsync

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

type AuthenticatedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthenticatedRequest) Reset() {
	*x = AuthenticatedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arsync_arsync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedRequest) ProtoMessage() {}

func (x *AuthenticatedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_arsync_arsync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedRequest.ProtoReflect.Descriptor instead.
func (*AuthenticatedRequest) Descriptor() ([]byte, []int) {
	return file_arsync_arsync_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticatedRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticatedRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PrepareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string                `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Auth *AuthenticatedRequest `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *PrepareRequest) Reset() {
	*x = PrepareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arsync_arsync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareRequest) ProtoMessage() {}

func (x *PrepareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_arsync_arsync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareRequest.ProtoReflect.Descriptor instead.
func (*PrepareRequest) Descriptor() ([]byte, []int) {
	return file_arsync_arsync_proto_rawDescGZIP(), []int{1}
}

func (x *PrepareRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PrepareRequest) GetAuth() *AuthenticatedRequest {
	if x != nil {
		return x.Auth
	}
	return nil
}

type PrepareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PrepareResponse) Reset() {
	*x = PrepareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arsync_arsync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareResponse) ProtoMessage() {}

func (x *PrepareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arsync_arsync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareResponse.ProtoReflect.Descriptor instead.
func (*PrepareResponse) Descriptor() ([]byte, []int) {
	return file_arsync_arsync_proto_rawDescGZIP(), []int{2}
}

func (x *PrepareResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arsync_arsync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arsync_arsync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_arsync_arsync_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_arsync_arsync_proto protoreflect.FileDescriptor

var file_arsync_arsync_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x61, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x22, 0x4e, 0x0a,
	0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x56, 0x0a,
	0x0e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x2b, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x24, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x32, 0x46, 0x0a, 0x06, 0x41, 0x72, 0x73, 0x79,
	0x6e, 0x63, 0x12, 0x3c, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x16, 0x2e,
	0x61, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a,
	0x61, 0x67, 0x67, 0x79, 0x2f, 0x61, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x61, 0x72, 0x73, 0x79,
	0x6e, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arsync_arsync_proto_rawDescOnce sync.Once
	file_arsync_arsync_proto_rawDescData = file_arsync_arsync_proto_rawDesc
)

func file_arsync_arsync_proto_rawDescGZIP() []byte {
	file_arsync_arsync_proto_rawDescOnce.Do(func() {
		file_arsync_arsync_proto_rawDescData = protoimpl.X.CompressGZIP(file_arsync_arsync_proto_rawDescData)
	})
	return file_arsync_arsync_proto_rawDescData
}

var file_arsync_arsync_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_arsync_arsync_proto_goTypes = []interface{}{
	(*AuthenticatedRequest)(nil), // 0: arsync.AuthenticatedRequest
	(*PrepareRequest)(nil),       // 1: arsync.PrepareRequest
	(*PrepareResponse)(nil),      // 2: arsync.PrepareResponse
	(*ListResponse)(nil),         // 3: arsync.ListResponse
}
var file_arsync_arsync_proto_depIdxs = []int32{
	0, // 0: arsync.PrepareRequest.auth:type_name -> arsync.AuthenticatedRequest
	1, // 1: arsync.Arsync.Prepare:input_type -> arsync.PrepareRequest
	2, // 2: arsync.Arsync.Prepare:output_type -> arsync.PrepareResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_arsync_arsync_proto_init() }
func file_arsync_arsync_proto_init() {
	if File_arsync_arsync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arsync_arsync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatedRequest); i {
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
		file_arsync_arsync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareRequest); i {
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
		file_arsync_arsync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareResponse); i {
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
		file_arsync_arsync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_arsync_arsync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arsync_arsync_proto_goTypes,
		DependencyIndexes: file_arsync_arsync_proto_depIdxs,
		MessageInfos:      file_arsync_arsync_proto_msgTypes,
	}.Build()
	File_arsync_arsync_proto = out.File
	file_arsync_arsync_proto_rawDesc = nil
	file_arsync_arsync_proto_goTypes = nil
	file_arsync_arsync_proto_depIdxs = nil
}
