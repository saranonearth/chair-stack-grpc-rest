// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: addGarment.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Garment string `protobuf:"bytes,2,opt,name=Garment,proto3" json:"Garment,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addGarment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_addGarment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_addGarment_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Item) GetGarment() string {
	if x != nil {
		return x.Garment
	}
	return ""
}

type ItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Item `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *ItemList) Reset() {
	*x = ItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addGarment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemList) ProtoMessage() {}

func (x *ItemList) ProtoReflect() protoreflect.Message {
	mi := &file_addGarment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemList.ProtoReflect.Descriptor instead.
func (*ItemList) Descriptor() ([]byte, []int) {
	return file_addGarment_proto_rawDescGZIP(), []int{1}
}

func (x *ItemList) GetList() []*Item {
	if x != nil {
		return x.List
	}
	return nil
}

var File_addGarment_proto protoreflect.FileDescriptor

var file_addGarment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x64, 0x47, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x04, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x47, 0x61, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x08, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x38, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x47, 0x61, 0x72,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_addGarment_proto_rawDescOnce sync.Once
	file_addGarment_proto_rawDescData = file_addGarment_proto_rawDesc
)

func file_addGarment_proto_rawDescGZIP() []byte {
	file_addGarment_proto_rawDescOnce.Do(func() {
		file_addGarment_proto_rawDescData = protoimpl.X.CompressGZIP(file_addGarment_proto_rawDescData)
	})
	return file_addGarment_proto_rawDescData
}

var file_addGarment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_addGarment_proto_goTypes = []interface{}{
	(*Item)(nil),     // 0: proto.Item
	(*ItemList)(nil), // 1: proto.ItemList
}
var file_addGarment_proto_depIdxs = []int32{
	0, // 0: proto.ItemList.List:type_name -> proto.Item
	0, // 1: proto.AddService.AddGarment:input_type -> proto.Item
	1, // 2: proto.AddService.AddGarment:output_type -> proto.ItemList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_addGarment_proto_init() }
func file_addGarment_proto_init() {
	if File_addGarment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_addGarment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_addGarment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemList); i {
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
			RawDescriptor: file_addGarment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_addGarment_proto_goTypes,
		DependencyIndexes: file_addGarment_proto_depIdxs,
		MessageInfos:      file_addGarment_proto_msgTypes,
	}.Build()
	File_addGarment_proto = out.File
	file_addGarment_proto_rawDesc = nil
	file_addGarment_proto_goTypes = nil
	file_addGarment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	AddGarment(ctx context.Context, in *Item, opts ...grpc.CallOption) (*ItemList, error)
}

type addServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddServiceClient(cc grpc.ClientConnInterface) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) AddGarment(ctx context.Context, in *Item, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/proto.AddService/AddGarment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	AddGarment(context.Context, *Item) (*ItemList, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) AddGarment(context.Context, *Item) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGarment not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_AddGarment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).AddGarment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/AddGarment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).AddGarment(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGarment",
			Handler:    _AddService_AddGarment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addGarment.proto",
}
