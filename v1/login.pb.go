// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: oauthapp/v1/login.proto

package oauthappv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
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

type GetLoginPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Continue string `protobuf:"bytes,1,opt,name=continue,proto3" json:"continue,omitempty"`
}

func (x *GetLoginPageRequest) Reset() {
	*x = GetLoginPageRequest{}
	mi := &file_oauthapp_v1_login_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLoginPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginPageRequest) ProtoMessage() {}

func (x *GetLoginPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthapp_v1_login_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginPageRequest.ProtoReflect.Descriptor instead.
func (*GetLoginPageRequest) Descriptor() ([]byte, []int) {
	return file_oauthapp_v1_login_proto_rawDescGZIP(), []int{0}
}

func (x *GetLoginPageRequest) GetContinue() string {
	if x != nil {
		return x.Continue
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pw string `protobuf:"bytes,2,opt,name=pw,proto3" json:"pw,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_oauthapp_v1_login_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthapp_v1_login_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_oauthapp_v1_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginRequest) GetPw() string {
	if x != nil {
		return x.Pw
	}
	return ""
}

type CallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *CallbackRequest) Reset() {
	*x = CallbackRequest{}
	mi := &file_oauthapp_v1_login_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackRequest) ProtoMessage() {}

func (x *CallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthapp_v1_login_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackRequest.ProtoReflect.Descriptor instead.
func (*CallbackRequest) Descriptor() ([]byte, []int) {
	return file_oauthapp_v1_login_proto_rawDescGZIP(), []int{2}
}

func (x *CallbackRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CallbackRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

var File_oauthapp_v1_login_proto protoreflect.FileDescriptor

var file_oauthapp_v1_login_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e,
	0x75, 0x65, 0x22, 0x2e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x70, 0x77, 0x22, 0x3b, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32,
	0xb3, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x56, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x20, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08,
	0x12, 0x06, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x4b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x19, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f,
	0x64, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x22, 0x06, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x32, 0x6c, 0x0a, 0x15, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x2f, 0x2a, 0x42, 0xa4, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x65, 0x74, 0x31, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70,
	0x70, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x17, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x4f, 0x61,
	0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_oauthapp_v1_login_proto_rawDescOnce sync.Once
	file_oauthapp_v1_login_proto_rawDescData = file_oauthapp_v1_login_proto_rawDesc
)

func file_oauthapp_v1_login_proto_rawDescGZIP() []byte {
	file_oauthapp_v1_login_proto_rawDescOnce.Do(func() {
		file_oauthapp_v1_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_oauthapp_v1_login_proto_rawDescData)
	})
	return file_oauthapp_v1_login_proto_rawDescData
}

var file_oauthapp_v1_login_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_oauthapp_v1_login_proto_goTypes = []any{
	(*GetLoginPageRequest)(nil), // 0: oauthapp.v1.GetLoginPageRequest
	(*LoginRequest)(nil),        // 1: oauthapp.v1.LoginRequest
	(*CallbackRequest)(nil),     // 2: oauthapp.v1.CallbackRequest
	(*httpbody.HttpBody)(nil),   // 3: google.api.HttpBody
}
var file_oauthapp_v1_login_proto_depIdxs = []int32{
	0, // 0: oauthapp.v1.LoginService.GetLoginPage:input_type -> oauthapp.v1.GetLoginPageRequest
	1, // 1: oauthapp.v1.LoginService.Login:input_type -> oauthapp.v1.LoginRequest
	2, // 2: oauthapp.v1.SocialCallbackService.Callback:input_type -> oauthapp.v1.CallbackRequest
	3, // 3: oauthapp.v1.LoginService.GetLoginPage:output_type -> google.api.HttpBody
	3, // 4: oauthapp.v1.LoginService.Login:output_type -> google.api.HttpBody
	3, // 5: oauthapp.v1.SocialCallbackService.Callback:output_type -> google.api.HttpBody
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oauthapp_v1_login_proto_init() }
func file_oauthapp_v1_login_proto_init() {
	if File_oauthapp_v1_login_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oauthapp_v1_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_oauthapp_v1_login_proto_goTypes,
		DependencyIndexes: file_oauthapp_v1_login_proto_depIdxs,
		MessageInfos:      file_oauthapp_v1_login_proto_msgTypes,
	}.Build()
	File_oauthapp_v1_login_proto = out.File
	file_oauthapp_v1_login_proto_rawDesc = nil
	file_oauthapp_v1_login_proto_goTypes = nil
	file_oauthapp_v1_login_proto_depIdxs = nil
}
