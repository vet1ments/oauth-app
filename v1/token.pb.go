// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: oauthapp/v1/token.proto

package oauthappv1

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

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	mi := &file_oauthapp_v1_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthapp_v1_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_oauthapp_v1_token_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type GetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken           string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiresIn             int64  `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	RefreshToken          string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn int64  `protobuf:"varint,4,opt,name=refresh_token_expires_in,json=refreshTokenExpiresIn,proto3" json:"refresh_token_expires_in,omitempty"`
}

func (x *GetTokenResponse) Reset() {
	*x = GetTokenResponse{}
	mi := &file_oauthapp_v1_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenResponse) ProtoMessage() {}

func (x *GetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauthapp_v1_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenResponse.ProtoReflect.Descriptor instead.
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return file_oauthapp_v1_token_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GetTokenResponse) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *GetTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *GetTokenResponse) GetRefreshTokenExpiresIn() int64 {
	if x != nil {
		return x.RefreshTokenExpiresIn
	}
	return 0
}

type GetTokenInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenInfoRequest) Reset() {
	*x = GetTokenInfoRequest{}
	mi := &file_oauthapp_v1_token_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTokenInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenInfoRequest) ProtoMessage() {}

func (x *GetTokenInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthapp_v1_token_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTokenInfoRequest) Descriptor() ([]byte, []int) {
	return file_oauthapp_v1_token_proto_rawDescGZIP(), []int{2}
}

func (x *GetTokenInfoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetTokenInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenInfoResponse) Reset() {
	*x = GetTokenInfoResponse{}
	mi := &file_oauthapp_v1_token_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTokenInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenInfoResponse) ProtoMessage() {}

func (x *GetTokenInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauthapp_v1_token_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenInfoResponse.ProtoReflect.Descriptor instead.
func (*GetTokenInfoResponse) Descriptor() ([]byte, []int) {
	return file_oauthapp_v1_token_proto_rawDescGZIP(), []int{3}
}

func (x *GetTokenInfoResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_oauthapp_v1_token_proto protoreflect.FileDescriptor

var file_oauthapp_v1_token_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22,
	0xb2, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x18, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x49, 0x6e, 0x22, 0x2b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32,
	0xab, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa4, 0x01,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x2e, 0x76,
	0x31, 0x42, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x74, 0x31,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x67, 0x6f, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa,
	0x02, 0x0b, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b,
	0x4f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x4f, 0x61,
	0x75, 0x74, 0x68, 0x61, 0x70, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x61, 0x70, 0x70,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oauthapp_v1_token_proto_rawDescOnce sync.Once
	file_oauthapp_v1_token_proto_rawDescData = file_oauthapp_v1_token_proto_rawDesc
)

func file_oauthapp_v1_token_proto_rawDescGZIP() []byte {
	file_oauthapp_v1_token_proto_rawDescOnce.Do(func() {
		file_oauthapp_v1_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_oauthapp_v1_token_proto_rawDescData)
	})
	return file_oauthapp_v1_token_proto_rawDescData
}

var file_oauthapp_v1_token_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_oauthapp_v1_token_proto_goTypes = []any{
	(*GetTokenRequest)(nil),      // 0: oauthapp.v1.GetTokenRequest
	(*GetTokenResponse)(nil),     // 1: oauthapp.v1.GetTokenResponse
	(*GetTokenInfoRequest)(nil),  // 2: oauthapp.v1.GetTokenInfoRequest
	(*GetTokenInfoResponse)(nil), // 3: oauthapp.v1.GetTokenInfoResponse
}
var file_oauthapp_v1_token_proto_depIdxs = []int32{
	0, // 0: oauthapp.v1.TokeService.GetToken:input_type -> oauthapp.v1.GetTokenRequest
	2, // 1: oauthapp.v1.TokeService.GetTokenInfo:input_type -> oauthapp.v1.GetTokenInfoRequest
	1, // 2: oauthapp.v1.TokeService.GetToken:output_type -> oauthapp.v1.GetTokenResponse
	3, // 3: oauthapp.v1.TokeService.GetTokenInfo:output_type -> oauthapp.v1.GetTokenInfoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oauthapp_v1_token_proto_init() }
func file_oauthapp_v1_token_proto_init() {
	if File_oauthapp_v1_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oauthapp_v1_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oauthapp_v1_token_proto_goTypes,
		DependencyIndexes: file_oauthapp_v1_token_proto_depIdxs,
		MessageInfos:      file_oauthapp_v1_token_proto_msgTypes,
	}.Build()
	File_oauthapp_v1_token_proto = out.File
	file_oauthapp_v1_token_proto_rawDesc = nil
	file_oauthapp_v1_token_proto_goTypes = nil
	file_oauthapp_v1_token_proto_depIdxs = nil
}
