// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.6
// source: Token.proto

package pb_github_com_golang_protobuf_TokenService

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

type UserClaim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Audience  string `protobuf:"bytes,2,opt,name=Audience,proto3" json:"Audience,omitempty"`
	ExpiresAt int64  `protobuf:"varint,3,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
	Id        string `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	IssuedAt  int64  `protobuf:"varint,5,opt,name=IssuedAt,proto3" json:"IssuedAt,omitempty"`
	Issuer    string `protobuf:"bytes,6,opt,name=Issuer,proto3" json:"Issuer,omitempty"`
	NotBefore int64  `protobuf:"varint,7,opt,name=NotBefore,proto3" json:"NotBefore,omitempty"`
	Subject   string `protobuf:"bytes,8,opt,name=Subject,proto3" json:"Subject,omitempty"`
}

func (x *UserClaim) Reset() {
	*x = UserClaim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserClaim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserClaim) ProtoMessage() {}

func (x *UserClaim) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserClaim.ProtoReflect.Descriptor instead.
func (*UserClaim) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{0}
}

func (x *UserClaim) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserClaim) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *UserClaim) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *UserClaim) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserClaim) GetIssuedAt() int64 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

func (x *UserClaim) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *UserClaim) GetNotBefore() int64 {
	if x != nil {
		return x.NotBefore
	}
	return 0
}

func (x *UserClaim) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_Token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_Token_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_Token_proto protoreflect.FileDescriptor

var file_Token_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xd5, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x64, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x22, 0x00, 0x42, 0x2f,
	0x5a, 0x2d, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x63,
	0x6f, 0x6d, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Token_proto_rawDescOnce sync.Once
	file_Token_proto_rawDescData = file_Token_proto_rawDesc
)

func file_Token_proto_rawDescGZIP() []byte {
	file_Token_proto_rawDescOnce.Do(func() {
		file_Token_proto_rawDescData = protoimpl.X.CompressGZIP(file_Token_proto_rawDescData)
	})
	return file_Token_proto_rawDescData
}

var file_Token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Token_proto_goTypes = []interface{}{
	(*UserClaim)(nil), // 0: pb.UserClaim
	(*Token)(nil),     // 1: pb.Token
}
var file_Token_proto_depIdxs = []int32{
	0, // 0: pb.TokenService.CreateToken:input_type -> pb.UserClaim
	1, // 1: pb.TokenService.ParserToken:input_type -> pb.Token
	1, // 2: pb.TokenService.CreateToken:output_type -> pb.Token
	0, // 3: pb.TokenService.ParserToken:output_type -> pb.UserClaim
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Token_proto_init() }
func file_Token_proto_init() {
	if File_Token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserClaim); i {
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
		file_Token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_Token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Token_proto_goTypes,
		DependencyIndexes: file_Token_proto_depIdxs,
		MessageInfos:      file_Token_proto_msgTypes,
	}.Build()
	File_Token_proto = out.File
	file_Token_proto_rawDesc = nil
	file_Token_proto_goTypes = nil
	file_Token_proto_depIdxs = nil
}
