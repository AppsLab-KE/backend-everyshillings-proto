// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: otp.proto

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

type CreateAndSendOtpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *CreateAndSendOtpReq) Reset() {
	*x = CreateAndSendOtpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAndSendOtpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAndSendOtpReq) ProtoMessage() {}

func (x *CreateAndSendOtpReq) ProtoReflect() protoreflect.Message {
	mi := &file_otp_proto_msgTypes[0]
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
	return file_otp_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAndSendOtpReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type CreateAndSendOtpRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	TrackingUuid string `protobuf:"bytes,3,opt,name=tracking_uuid,json=trackingUuid,proto3" json:"tracking_uuid,omitempty"`
}

func (x *CreateAndSendOtpRes) Reset() {
	*x = CreateAndSendOtpRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAndSendOtpRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAndSendOtpRes) ProtoMessage() {}

func (x *CreateAndSendOtpRes) ProtoReflect() protoreflect.Message {
	mi := &file_otp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAndSendOtpRes.ProtoReflect.Descriptor instead.
func (*CreateAndSendOtpRes) Descriptor() ([]byte, []int) {
	return file_otp_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAndSendOtpRes) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreateAndSendOtpRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateAndSendOtpRes) GetTrackingUuid() string {
	if x != nil {
		return x.TrackingUuid
	}
	return ""
}

type VerifyOTPReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtpCode      string `protobuf:"bytes,1,opt,name=otp_code,json=otpCode,proto3" json:"otp_code,omitempty"`
	TrackingUuid string `protobuf:"bytes,2,opt,name=tracking_uuid,json=trackingUuid,proto3" json:"tracking_uuid,omitempty"`
}

func (x *VerifyOTPReq) Reset() {
	*x = VerifyOTPReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPReq) ProtoMessage() {}

func (x *VerifyOTPReq) ProtoReflect() protoreflect.Message {
	mi := &file_otp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPReq.ProtoReflect.Descriptor instead.
func (*VerifyOTPReq) Descriptor() ([]byte, []int) {
	return file_otp_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyOTPReq) GetOtpCode() string {
	if x != nil {
		return x.OtpCode
	}
	return ""
}

func (x *VerifyOTPReq) GetTrackingUuid() string {
	if x != nil {
		return x.TrackingUuid
	}
	return ""
}

type VerifyOTPRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VerifyOTPRes) Reset() {
	*x = VerifyOTPRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPRes) ProtoMessage() {}

func (x *VerifyOTPRes) ProtoReflect() protoreflect.Message {
	mi := &file_otp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPRes.ProtoReflect.Descriptor instead.
func (*VerifyOTPRes) Descriptor() ([]byte, []int) {
	return file_otp_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyOTPRes) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *VerifyOTPRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResendOTPReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingId string `protobuf:"bytes,1,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
}

func (x *ResendOTPReq) Reset() {
	*x = ResendOTPReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendOTPReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendOTPReq) ProtoMessage() {}

func (x *ResendOTPReq) ProtoReflect() protoreflect.Message {
	mi := &file_otp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendOTPReq.ProtoReflect.Descriptor instead.
func (*ResendOTPReq) Descriptor() ([]byte, []int) {
	return file_otp_proto_rawDescGZIP(), []int{4}
}

func (x *ResendOTPReq) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

type ResendOTPRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	TrackingUuid string `protobuf:"bytes,3,opt,name=tracking_uuid,json=trackingUuid,proto3" json:"tracking_uuid,omitempty"`
}

func (x *ResendOTPRes) Reset() {
	*x = ResendOTPRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendOTPRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendOTPRes) ProtoMessage() {}

func (x *ResendOTPRes) ProtoReflect() protoreflect.Message {
	mi := &file_otp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendOTPRes.ProtoReflect.Descriptor instead.
func (*ResendOTPRes) Descriptor() ([]byte, []int) {
	return file_otp_proto_rawDescGZIP(), []int{5}
}

func (x *ResendOTPRes) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ResendOTPRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResendOTPRes) GetTrackingUuid() string {
	if x != nil {
		return x.TrackingUuid
	}
	return ""
}

var File_otp_proto protoreflect.FileDescriptor

var file_otp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6f, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x74, 0x70,
	0x22, 0x38, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6e,
	0x64, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x75, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x52, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x75, 0x69,
	0x64, 0x22, 0x4e, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x75, 0x69,
	0x64, 0x22, 0x49, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x6e, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x75, 0x69, 0x64, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x70, 0x70, 0x73,
	0x4c, 0x61, 0x62, 0x2d, 0x4b, 0x45, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x65,
	0x76, 0x65, 0x72, 0x79, 0x73, 0x68, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x6f, 0x74,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_otp_proto_rawDescOnce sync.Once
	file_otp_proto_rawDescData = file_otp_proto_rawDesc
)

func file_otp_proto_rawDescGZIP() []byte {
	file_otp_proto_rawDescOnce.Do(func() {
		file_otp_proto_rawDescData = protoimpl.X.CompressGZIP(file_otp_proto_rawDescData)
	})
	return file_otp_proto_rawDescData
}

var file_otp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_otp_proto_goTypes = []interface{}{
	(*CreateAndSendOtpReq)(nil), // 0: otp.CreateAndSendOtpReq
	(*CreateAndSendOtpRes)(nil), // 1: otp.CreateAndSendOtpRes
	(*VerifyOTPReq)(nil),        // 2: otp.VerifyOTPReq
	(*VerifyOTPRes)(nil),        // 3: otp.VerifyOTPRes
	(*ResendOTPReq)(nil),        // 4: otp.ResendOTPReq
	(*ResendOTPRes)(nil),        // 5: otp.ResendOTPRes
}
var file_otp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_otp_proto_init() }
func file_otp_proto_init() {
	if File_otp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_otp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_otp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAndSendOtpRes); i {
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
		file_otp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOTPReq); i {
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
		file_otp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOTPRes); i {
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
		file_otp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendOTPReq); i {
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
		file_otp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendOTPRes); i {
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
			RawDescriptor: file_otp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_otp_proto_goTypes,
		DependencyIndexes: file_otp_proto_depIdxs,
		MessageInfos:      file_otp_proto_msgTypes,
	}.Build()
	File_otp_proto = out.File
	file_otp_proto_rawDesc = nil
	file_otp_proto_goTypes = nil
	file_otp_proto_depIdxs = nil
}
