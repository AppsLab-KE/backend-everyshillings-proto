// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: db/server.proto

package db

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

type DefaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DefaultRequest) Reset() {
	*x = DefaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultRequest) ProtoMessage() {}

func (x *DefaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultRequest.ProtoReflect.Descriptor instead.
func (*DefaultRequest) Descriptor() ([]byte, []int) {
	return file_db_server_proto_rawDescGZIP(), []int{0}
}

type GetResourceByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *GetResourceByIdRequest) Reset() {
	*x = GetResourceByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByIdRequest) ProtoMessage() {}

func (x *GetResourceByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByIdRequest.ProtoReflect.Descriptor instead.
func (*GetResourceByIdRequest) Descriptor() ([]byte, []int) {
	return file_db_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetResourceByIdRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type KeyValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValueRequest) Reset() {
	*x = KeyValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueRequest) ProtoMessage() {}

func (x *KeyValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueRequest.ProtoReflect.Descriptor instead.
func (*KeyValueRequest) Descriptor() ([]byte, []int) {
	return file_db_server_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MultipleKeyValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*KeyValueRequest `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MultipleKeyValueRequest) Reset() {
	*x = MultipleKeyValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleKeyValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleKeyValueRequest) ProtoMessage() {}

func (x *MultipleKeyValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleKeyValueRequest.ProtoReflect.Descriptor instead.
func (*MultipleKeyValueRequest) Descriptor() ([]byte, []int) {
	return file_db_server_proto_rawDescGZIP(), []int{3}
}

func (x *MultipleKeyValueRequest) GetData() []*KeyValueRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

type DefaultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error   bool   `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DefaultResponse) Reset() {
	*x = DefaultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultResponse) ProtoMessage() {}

func (x *DefaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultResponse.ProtoReflect.Descriptor instead.
func (*DefaultResponse) Descriptor() ([]byte, []int) {
	return file_db_server_proto_rawDescGZIP(), []int{4}
}

func (x *DefaultResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DefaultResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_db_server_proto_rawDescGZIP(), []int{5}
}

func (x *HealthResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_db_server_proto protoreflect.FileDescriptor

var file_db_server_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x64, 0x62, 0x1a, 0x0d, 0x64, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x39, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x17,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x41, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x9e, 0x01, 0x0a, 0x09, 0x44, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x64,
	0x62, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x64, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x11, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e, 0x64, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x13, 0x2e, 0x64, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x64, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x6c, 0x61, 0x62, 0x2f, 0x74, 0x77, 0x61, 0x6c, 0x61, 0x2f, 0x64, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_server_proto_rawDescOnce sync.Once
	file_db_server_proto_rawDescData = file_db_server_proto_rawDesc
)

func file_db_server_proto_rawDescGZIP() []byte {
	file_db_server_proto_rawDescOnce.Do(func() {
		file_db_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_server_proto_rawDescData)
	})
	return file_db_server_proto_rawDescData
}

var file_db_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_db_server_proto_goTypes = []interface{}{
	(*DefaultRequest)(nil),          // 0: db.DefaultRequest
	(*GetResourceByIdRequest)(nil),  // 1: db.GetResourceByIdRequest
	(*KeyValueRequest)(nil),         // 2: db.KeyValueRequest
	(*MultipleKeyValueRequest)(nil), // 3: db.MultipleKeyValueRequest
	(*DefaultResponse)(nil),         // 4: db.DefaultResponse
	(*HealthResponse)(nil),          // 5: db.HealthResponse
	(*CreateUserReq)(nil),           // 6: db.CreateUserReq
	(*User)(nil),                    // 7: db.User
}
var file_db_server_proto_depIdxs = []int32{
	2, // 0: db.MultipleKeyValueRequest.data:type_name -> db.KeyValueRequest
	0, // 1: db.DbService.HealthCheck:input_type -> db.DefaultRequest
	6, // 2: db.DbService.CreateUser:input_type -> db.CreateUserReq
	2, // 3: db.DbService.GetUserByField:input_type -> db.KeyValueRequest
	5, // 4: db.DbService.HealthCheck:output_type -> db.HealthResponse
	7, // 5: db.DbService.CreateUser:output_type -> db.User
	7, // 6: db.DbService.GetUserByField:output_type -> db.User
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_db_server_proto_init() }
func file_db_server_proto_init() {
	if File_db_server_proto != nil {
		return
	}
	file_db_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_db_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultRequest); i {
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
		file_db_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByIdRequest); i {
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
		file_db_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueRequest); i {
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
		file_db_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleKeyValueRequest); i {
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
		file_db_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultResponse); i {
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
		file_db_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
			RawDescriptor: file_db_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_db_server_proto_goTypes,
		DependencyIndexes: file_db_server_proto_depIdxs,
		MessageInfos:      file_db_server_proto_msgTypes,
	}.Build()
	File_db_server_proto = out.File
	file_db_server_proto_rawDesc = nil
	file_db_server_proto_goTypes = nil
	file_db_server_proto_depIdxs = nil
}
