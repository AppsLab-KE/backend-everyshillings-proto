// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: otp/server.proto

package otp

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
		mi := &file_otp_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultRequest) ProtoMessage() {}

func (x *DefaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otp_server_proto_msgTypes[0]
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
	return file_otp_server_proto_rawDescGZIP(), []int{0}
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
		mi := &file_otp_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otp_server_proto_msgTypes[1]
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
	return file_otp_server_proto_rawDescGZIP(), []int{1}
}

func (x *HealthResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
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
		mi := &file_otp_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultResponse) ProtoMessage() {}

func (x *DefaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otp_server_proto_msgTypes[2]
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
	return file_otp_server_proto_rawDescGZIP(), []int{2}
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

type CreateAndSendOtpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	OtpLength   string `protobuf:"bytes,2,opt,name=otp_length,json=otpLength,proto3" json:"otp_length,omitempty"`
}

func (x *CreateAndSendOtpReq) Reset() {
	*x = CreateAndSendOtpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAndSendOtpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAndSendOtpReq) ProtoMessage() {}

func (x *CreateAndSendOtpReq) ProtoReflect() protoreflect.Message {
	mi := &file_otp_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAndSendOtpReq.ProtoReflect.Descriptor instead.
func (*CreateAndSendOtpReq) Descriptor() ([]byte, []int) {
	return file_otp_server_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAndSendOtpReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateAndSendOtpReq) GetOtpLength() string {
	if x != nil {
		return x.OtpLength
	}
	return ""
}

var File_otp_server_proto protoreflect.FileDescriptor

var file_otp_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x74, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x57, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x74, 0x70, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x74, 0x70, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x32, 0x89, 0x01, 0x0a, 0x0a, 0x4f, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x13, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x12, 0x18, 0x2e,
	0x6f, 0x74, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6e,
	0x64, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x70, 0x70, 0x73,
	0x4c, 0x61, 0x62, 0x2d, 0x4b, 0x45, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x65,
	0x76, 0x65, 0x72, 0x79, 0x73, 0x68, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x6f, 0x74,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_otp_server_proto_rawDescOnce sync.Once
	file_otp_server_proto_rawDescData = file_otp_server_proto_rawDesc
)

func file_otp_server_proto_rawDescGZIP() []byte {
	file_otp_server_proto_rawDescOnce.Do(func() {
		file_otp_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_otp_server_proto_rawDescData)
	})
	return file_otp_server_proto_rawDescData
}

var file_otp_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_otp_server_proto_goTypes = []interface{}{
	(*DefaultRequest)(nil),      // 0: otp.DefaultRequest
	(*HealthResponse)(nil),      // 1: otp.HealthResponse
	(*DefaultResponse)(nil),     // 2: otp.DefaultResponse
	(*CreateAndSendOtpReq)(nil), // 3: otp.CreateAndSendOtpReq
}
var file_otp_server_proto_depIdxs = []int32{
	0, // 0: otp.OtpService.HealthCheck:input_type -> otp.DefaultRequest
	3, // 1: otp.OtpService.CreateAndSendOtp:input_type -> otp.CreateAndSendOtpReq
	1, // 2: otp.OtpService.HealthCheck:output_type -> otp.HealthResponse
	2, // 3: otp.OtpService.CreateAndSendOtp:output_type -> otp.DefaultResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_otp_server_proto_init() }
func file_otp_server_proto_init() {
	if File_otp_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_otp_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_otp_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_otp_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_otp_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAndSendOtpReq); i {
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
			RawDescriptor: file_otp_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_otp_server_proto_goTypes,
		DependencyIndexes: file_otp_server_proto_depIdxs,
		MessageInfos:      file_otp_server_proto_msgTypes,
	}.Build()
	File_otp_server_proto = out.File
	file_otp_server_proto_rawDesc = nil
	file_otp_server_proto_goTypes = nil
	file_otp_server_proto_depIdxs = nil
}
