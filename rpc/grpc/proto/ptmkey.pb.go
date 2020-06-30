// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.7.1
// source: ptmkey.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	types "github.com/qlcchain/go-qlc/rpc/grpc/proto/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PtmKeyUpdateParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Btype   string `protobuf:"bytes,2,opt,name=btype,proto3" json:"btype,omitempty"`
	Pubkey  string `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (x *PtmKeyUpdateParam) Reset() {
	*x = PtmKeyUpdateParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptmkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PtmKeyUpdateParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PtmKeyUpdateParam) ProtoMessage() {}

func (x *PtmKeyUpdateParam) ProtoReflect() protoreflect.Message {
	mi := &file_ptmkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PtmKeyUpdateParam.ProtoReflect.Descriptor instead.
func (*PtmKeyUpdateParam) Descriptor() ([]byte, []int) {
	return file_ptmkey_proto_rawDescGZIP(), []int{0}
}

func (x *PtmKeyUpdateParam) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PtmKeyUpdateParam) GetBtype() string {
	if x != nil {
		return x.Btype
	}
	return ""
}

func (x *PtmKeyUpdateParam) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

type PtmKeyUpdateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params []*PtmKeyUpdateParam `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *PtmKeyUpdateParams) Reset() {
	*x = PtmKeyUpdateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptmkey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PtmKeyUpdateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PtmKeyUpdateParams) ProtoMessage() {}

func (x *PtmKeyUpdateParams) ProtoReflect() protoreflect.Message {
	mi := &file_ptmkey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PtmKeyUpdateParams.ProtoReflect.Descriptor instead.
func (*PtmKeyUpdateParams) Descriptor() ([]byte, []int) {
	return file_ptmkey_proto_rawDescGZIP(), []int{1}
}

func (x *PtmKeyUpdateParams) GetParams() []*PtmKeyUpdateParam {
	if x != nil {
		return x.Params
	}
	return nil
}

type PtmKeyByAccountAndBtypeParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Btype   string `protobuf:"bytes,2,opt,name=btype,proto3" json:"btype,omitempty"`
}

func (x *PtmKeyByAccountAndBtypeParam) Reset() {
	*x = PtmKeyByAccountAndBtypeParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptmkey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PtmKeyByAccountAndBtypeParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PtmKeyByAccountAndBtypeParam) ProtoMessage() {}

func (x *PtmKeyByAccountAndBtypeParam) ProtoReflect() protoreflect.Message {
	mi := &file_ptmkey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PtmKeyByAccountAndBtypeParam.ProtoReflect.Descriptor instead.
func (*PtmKeyByAccountAndBtypeParam) Descriptor() ([]byte, []int) {
	return file_ptmkey_proto_rawDescGZIP(), []int{2}
}

func (x *PtmKeyByAccountAndBtypeParam) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PtmKeyByAccountAndBtypeParam) GetBtype() string {
	if x != nil {
		return x.Btype
	}
	return ""
}

var File_ptmkey_proto protoreflect.FileDescriptor

var file_ptmkey_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x74, 0x6d, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x11, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x22, 0x46, 0x0a, 0x12, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4e, 0x0a, 0x1c, 0x50,
	0x74, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6e,
	0x64, 0x42, 0x74, 0x79, 0x70, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x74, 0x79, 0x70, 0x65, 0x32, 0xdc, 0x03, 0x0a, 0x09,
	0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x41, 0x50, 0x49, 0x12, 0x69, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x11, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x70, 0x74, 0x6d, 0x6b, 0x65, 0x79, 0x2f,
	0x67, 0x65, 0x74, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x74, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x74, 0x6d, 0x4b, 0x65,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x42, 0x74, 0x79, 0x70, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x70,
	0x74, 0x6d, 0x6b, 0x65, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x63, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x70, 0x74, 0x6d, 0x6b, 0x65, 0x79, 0x2f, 0x67, 0x65, 0x74,
	0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x88, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x42, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x42, 0x74, 0x79, 0x70, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x74, 0x6d, 0x4b,
	0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x70, 0x74, 0x6d, 0x6b, 0x65, 0x79, 0x2f,
	0x67, 0x65, 0x74, 0x50, 0x74, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x6e, 0x64, 0x42, 0x74, 0x79, 0x70, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ptmkey_proto_rawDescOnce sync.Once
	file_ptmkey_proto_rawDescData = file_ptmkey_proto_rawDesc
)

func file_ptmkey_proto_rawDescGZIP() []byte {
	file_ptmkey_proto_rawDescOnce.Do(func() {
		file_ptmkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_ptmkey_proto_rawDescData)
	})
	return file_ptmkey_proto_rawDescData
}

var file_ptmkey_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ptmkey_proto_goTypes = []interface{}{
	(*PtmKeyUpdateParam)(nil),            // 0: proto.PtmKeyUpdateParam
	(*PtmKeyUpdateParams)(nil),           // 1: proto.PtmKeyUpdateParams
	(*PtmKeyByAccountAndBtypeParam)(nil), // 2: proto.PtmKeyByAccountAndBtypeParam
	(*types.Address)(nil),                // 3: types.Address
	(*types.StateBlock)(nil),             // 4: types.StateBlock
}
var file_ptmkey_proto_depIdxs = []int32{
	0, // 0: proto.PtmKeyUpdateParams.params:type_name -> proto.PtmKeyUpdateParam
	0, // 1: proto.PtmKeyAPI.GetPtmKeyUpdateBlock:input_type -> proto.PtmKeyUpdateParam
	2, // 2: proto.PtmKeyAPI.GetPtmKeyDeleteBlock:input_type -> proto.PtmKeyByAccountAndBtypeParam
	3, // 3: proto.PtmKeyAPI.GetPtmKeyByAccount:input_type -> types.Address
	2, // 4: proto.PtmKeyAPI.GetPtmKeyByAccountAndBtype:input_type -> proto.PtmKeyByAccountAndBtypeParam
	4, // 5: proto.PtmKeyAPI.GetPtmKeyUpdateBlock:output_type -> types.StateBlock
	4, // 6: proto.PtmKeyAPI.GetPtmKeyDeleteBlock:output_type -> types.StateBlock
	1, // 7: proto.PtmKeyAPI.GetPtmKeyByAccount:output_type -> proto.PtmKeyUpdateParams
	1, // 8: proto.PtmKeyAPI.GetPtmKeyByAccountAndBtype:output_type -> proto.PtmKeyUpdateParams
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ptmkey_proto_init() }
func file_ptmkey_proto_init() {
	if File_ptmkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ptmkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PtmKeyUpdateParam); i {
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
		file_ptmkey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PtmKeyUpdateParams); i {
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
		file_ptmkey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PtmKeyByAccountAndBtypeParam); i {
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
			RawDescriptor: file_ptmkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ptmkey_proto_goTypes,
		DependencyIndexes: file_ptmkey_proto_depIdxs,
		MessageInfos:      file_ptmkey_proto_msgTypes,
	}.Build()
	File_ptmkey_proto = out.File
	file_ptmkey_proto_rawDesc = nil
	file_ptmkey_proto_goTypes = nil
	file_ptmkey_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PtmKeyAPIClient is the client API for PtmKeyAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PtmKeyAPIClient interface {
	GetPtmKeyUpdateBlock(ctx context.Context, in *PtmKeyUpdateParam, opts ...grpc.CallOption) (*types.StateBlock, error)
	GetPtmKeyDeleteBlock(ctx context.Context, in *PtmKeyByAccountAndBtypeParam, opts ...grpc.CallOption) (*types.StateBlock, error)
	GetPtmKeyByAccount(ctx context.Context, in *types.Address, opts ...grpc.CallOption) (*PtmKeyUpdateParams, error)
	GetPtmKeyByAccountAndBtype(ctx context.Context, in *PtmKeyByAccountAndBtypeParam, opts ...grpc.CallOption) (*PtmKeyUpdateParams, error)
}

type ptmKeyAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPtmKeyAPIClient(cc grpc.ClientConnInterface) PtmKeyAPIClient {
	return &ptmKeyAPIClient{cc}
}

func (c *ptmKeyAPIClient) GetPtmKeyUpdateBlock(ctx context.Context, in *PtmKeyUpdateParam, opts ...grpc.CallOption) (*types.StateBlock, error) {
	out := new(types.StateBlock)
	err := c.cc.Invoke(ctx, "/proto.PtmKeyAPI/GetPtmKeyUpdateBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ptmKeyAPIClient) GetPtmKeyDeleteBlock(ctx context.Context, in *PtmKeyByAccountAndBtypeParam, opts ...grpc.CallOption) (*types.StateBlock, error) {
	out := new(types.StateBlock)
	err := c.cc.Invoke(ctx, "/proto.PtmKeyAPI/GetPtmKeyDeleteBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ptmKeyAPIClient) GetPtmKeyByAccount(ctx context.Context, in *types.Address, opts ...grpc.CallOption) (*PtmKeyUpdateParams, error) {
	out := new(PtmKeyUpdateParams)
	err := c.cc.Invoke(ctx, "/proto.PtmKeyAPI/GetPtmKeyByAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ptmKeyAPIClient) GetPtmKeyByAccountAndBtype(ctx context.Context, in *PtmKeyByAccountAndBtypeParam, opts ...grpc.CallOption) (*PtmKeyUpdateParams, error) {
	out := new(PtmKeyUpdateParams)
	err := c.cc.Invoke(ctx, "/proto.PtmKeyAPI/GetPtmKeyByAccountAndBtype", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PtmKeyAPIServer is the server API for PtmKeyAPI service.
type PtmKeyAPIServer interface {
	GetPtmKeyUpdateBlock(context.Context, *PtmKeyUpdateParam) (*types.StateBlock, error)
	GetPtmKeyDeleteBlock(context.Context, *PtmKeyByAccountAndBtypeParam) (*types.StateBlock, error)
	GetPtmKeyByAccount(context.Context, *types.Address) (*PtmKeyUpdateParams, error)
	GetPtmKeyByAccountAndBtype(context.Context, *PtmKeyByAccountAndBtypeParam) (*PtmKeyUpdateParams, error)
}

// UnimplementedPtmKeyAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPtmKeyAPIServer struct {
}

func (*UnimplementedPtmKeyAPIServer) GetPtmKeyUpdateBlock(context.Context, *PtmKeyUpdateParam) (*types.StateBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPtmKeyUpdateBlock not implemented")
}
func (*UnimplementedPtmKeyAPIServer) GetPtmKeyDeleteBlock(context.Context, *PtmKeyByAccountAndBtypeParam) (*types.StateBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPtmKeyDeleteBlock not implemented")
}
func (*UnimplementedPtmKeyAPIServer) GetPtmKeyByAccount(context.Context, *types.Address) (*PtmKeyUpdateParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPtmKeyByAccount not implemented")
}
func (*UnimplementedPtmKeyAPIServer) GetPtmKeyByAccountAndBtype(context.Context, *PtmKeyByAccountAndBtypeParam) (*PtmKeyUpdateParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPtmKeyByAccountAndBtype not implemented")
}

func RegisterPtmKeyAPIServer(s *grpc.Server, srv PtmKeyAPIServer) {
	s.RegisterService(&_PtmKeyAPI_serviceDesc, srv)
}

func _PtmKeyAPI_GetPtmKeyUpdateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PtmKeyUpdateParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtmKeyAPIServer).GetPtmKeyUpdateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PtmKeyAPI/GetPtmKeyUpdateBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtmKeyAPIServer).GetPtmKeyUpdateBlock(ctx, req.(*PtmKeyUpdateParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _PtmKeyAPI_GetPtmKeyDeleteBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PtmKeyByAccountAndBtypeParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtmKeyAPIServer).GetPtmKeyDeleteBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PtmKeyAPI/GetPtmKeyDeleteBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtmKeyAPIServer).GetPtmKeyDeleteBlock(ctx, req.(*PtmKeyByAccountAndBtypeParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _PtmKeyAPI_GetPtmKeyByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtmKeyAPIServer).GetPtmKeyByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PtmKeyAPI/GetPtmKeyByAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtmKeyAPIServer).GetPtmKeyByAccount(ctx, req.(*types.Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _PtmKeyAPI_GetPtmKeyByAccountAndBtype_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PtmKeyByAccountAndBtypeParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PtmKeyAPIServer).GetPtmKeyByAccountAndBtype(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PtmKeyAPI/GetPtmKeyByAccountAndBtype",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PtmKeyAPIServer).GetPtmKeyByAccountAndBtype(ctx, req.(*PtmKeyByAccountAndBtypeParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _PtmKeyAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PtmKeyAPI",
	HandlerType: (*PtmKeyAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPtmKeyUpdateBlock",
			Handler:    _PtmKeyAPI_GetPtmKeyUpdateBlock_Handler,
		},
		{
			MethodName: "GetPtmKeyDeleteBlock",
			Handler:    _PtmKeyAPI_GetPtmKeyDeleteBlock_Handler,
		},
		{
			MethodName: "GetPtmKeyByAccount",
			Handler:    _PtmKeyAPI_GetPtmKeyByAccount_Handler,
		},
		{
			MethodName: "GetPtmKeyByAccountAndBtype",
			Handler:    _PtmKeyAPI_GetPtmKeyByAccountAndBtype_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ptmkey.proto",
}