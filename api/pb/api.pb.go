// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NewStoreRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewStoreRequest) Reset()         { *m = NewStoreRequest{} }
func (m *NewStoreRequest) String() string { return proto.CompactTextString(m) }
func (*NewStoreRequest) ProtoMessage()    {}
func (*NewStoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_120f3cfc08b047b0, []int{0}
}
func (m *NewStoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewStoreRequest.Unmarshal(m, b)
}
func (m *NewStoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewStoreRequest.Marshal(b, m, deterministic)
}
func (dst *NewStoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewStoreRequest.Merge(dst, src)
}
func (m *NewStoreRequest) XXX_Size() int {
	return xxx_messageInfo_NewStoreRequest.Size(m)
}
func (m *NewStoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewStoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewStoreRequest proto.InternalMessageInfo

type NewStoreReply struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewStoreReply) Reset()         { *m = NewStoreReply{} }
func (m *NewStoreReply) String() string { return proto.CompactTextString(m) }
func (*NewStoreReply) ProtoMessage()    {}
func (*NewStoreReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_120f3cfc08b047b0, []int{1}
}
func (m *NewStoreReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewStoreReply.Unmarshal(m, b)
}
func (m *NewStoreReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewStoreReply.Marshal(b, m, deterministic)
}
func (dst *NewStoreReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewStoreReply.Merge(dst, src)
}
func (m *NewStoreReply) XXX_Size() int {
	return xxx_messageInfo_NewStoreReply.Size(m)
}
func (m *NewStoreReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NewStoreReply.DiscardUnknown(m)
}

var xxx_messageInfo_NewStoreReply proto.InternalMessageInfo

func (m *NewStoreReply) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RegisterSchemaRequest struct {
	StoreID              string   `protobuf:"bytes,1,opt,name=storeID,proto3" json:"storeID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Schema               string   `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterSchemaRequest) Reset()         { *m = RegisterSchemaRequest{} }
func (m *RegisterSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterSchemaRequest) ProtoMessage()    {}
func (*RegisterSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_120f3cfc08b047b0, []int{2}
}
func (m *RegisterSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterSchemaRequest.Unmarshal(m, b)
}
func (m *RegisterSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterSchemaRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterSchemaRequest.Merge(dst, src)
}
func (m *RegisterSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterSchemaRequest.Size(m)
}
func (m *RegisterSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterSchemaRequest proto.InternalMessageInfo

func (m *RegisterSchemaRequest) GetStoreID() string {
	if m != nil {
		return m.StoreID
	}
	return ""
}

func (m *RegisterSchemaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterSchemaRequest) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

type RegisterSchemaReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterSchemaReply) Reset()         { *m = RegisterSchemaReply{} }
func (m *RegisterSchemaReply) String() string { return proto.CompactTextString(m) }
func (*RegisterSchemaReply) ProtoMessage()    {}
func (*RegisterSchemaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_120f3cfc08b047b0, []int{3}
}
func (m *RegisterSchemaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterSchemaReply.Unmarshal(m, b)
}
func (m *RegisterSchemaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterSchemaReply.Marshal(b, m, deterministic)
}
func (dst *RegisterSchemaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterSchemaReply.Merge(dst, src)
}
func (m *RegisterSchemaReply) XXX_Size() int {
	return xxx_messageInfo_RegisterSchemaReply.Size(m)
}
func (m *RegisterSchemaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterSchemaReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterSchemaReply proto.InternalMessageInfo

type ModelCreateRequest struct {
	StoreID              string   `protobuf:"bytes,1,opt,name=storeID,proto3" json:"storeID,omitempty"`
	ModelName            string   `protobuf:"bytes,2,opt,name=modelName,proto3" json:"modelName,omitempty"`
	Values               []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelCreateRequest) Reset()         { *m = ModelCreateRequest{} }
func (m *ModelCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ModelCreateRequest) ProtoMessage()    {}
func (*ModelCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_120f3cfc08b047b0, []int{4}
}
func (m *ModelCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelCreateRequest.Unmarshal(m, b)
}
func (m *ModelCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelCreateRequest.Marshal(b, m, deterministic)
}
func (dst *ModelCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelCreateRequest.Merge(dst, src)
}
func (m *ModelCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ModelCreateRequest.Size(m)
}
func (m *ModelCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModelCreateRequest proto.InternalMessageInfo

func (m *ModelCreateRequest) GetStoreID() string {
	if m != nil {
		return m.StoreID
	}
	return ""
}

func (m *ModelCreateRequest) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *ModelCreateRequest) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ModelCreateReply struct {
	Entities             []string `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelCreateReply) Reset()         { *m = ModelCreateReply{} }
func (m *ModelCreateReply) String() string { return proto.CompactTextString(m) }
func (*ModelCreateReply) ProtoMessage()    {}
func (*ModelCreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_120f3cfc08b047b0, []int{5}
}
func (m *ModelCreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelCreateReply.Unmarshal(m, b)
}
func (m *ModelCreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelCreateReply.Marshal(b, m, deterministic)
}
func (dst *ModelCreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelCreateReply.Merge(dst, src)
}
func (m *ModelCreateReply) XXX_Size() int {
	return xxx_messageInfo_ModelCreateReply.Size(m)
}
func (m *ModelCreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelCreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_ModelCreateReply proto.InternalMessageInfo

func (m *ModelCreateReply) GetEntities() []string {
	if m != nil {
		return m.Entities
	}
	return nil
}

type ListenRequest struct {
	StoreID              string   `protobuf:"bytes,1,opt,name=storeID,proto3" json:"storeID,omitempty"`
	ModelName            string   `protobuf:"bytes,2,opt,name=modelName,proto3" json:"modelName,omitempty"`
	EntityID             string   `protobuf:"bytes,3,opt,name=entityID,proto3" json:"entityID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenRequest) Reset()         { *m = ListenRequest{} }
func (m *ListenRequest) String() string { return proto.CompactTextString(m) }
func (*ListenRequest) ProtoMessage()    {}
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_120f3cfc08b047b0, []int{6}
}
func (m *ListenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenRequest.Unmarshal(m, b)
}
func (m *ListenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenRequest.Marshal(b, m, deterministic)
}
func (dst *ListenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenRequest.Merge(dst, src)
}
func (m *ListenRequest) XXX_Size() int {
	return xxx_messageInfo_ListenRequest.Size(m)
}
func (m *ListenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenRequest proto.InternalMessageInfo

func (m *ListenRequest) GetStoreID() string {
	if m != nil {
		return m.StoreID
	}
	return ""
}

func (m *ListenRequest) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *ListenRequest) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

type ListenReply struct {
	Entity               string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenReply) Reset()         { *m = ListenReply{} }
func (m *ListenReply) String() string { return proto.CompactTextString(m) }
func (*ListenReply) ProtoMessage()    {}
func (*ListenReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_120f3cfc08b047b0, []int{7}
}
func (m *ListenReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenReply.Unmarshal(m, b)
}
func (m *ListenReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenReply.Marshal(b, m, deterministic)
}
func (dst *ListenReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenReply.Merge(dst, src)
}
func (m *ListenReply) XXX_Size() int {
	return xxx_messageInfo_ListenReply.Size(m)
}
func (m *ListenReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListenReply proto.InternalMessageInfo

func (m *ListenReply) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func init() {
	proto.RegisterType((*NewStoreRequest)(nil), "api.pb.NewStoreRequest")
	proto.RegisterType((*NewStoreReply)(nil), "api.pb.NewStoreReply")
	proto.RegisterType((*RegisterSchemaRequest)(nil), "api.pb.RegisterSchemaRequest")
	proto.RegisterType((*RegisterSchemaReply)(nil), "api.pb.RegisterSchemaReply")
	proto.RegisterType((*ModelCreateRequest)(nil), "api.pb.ModelCreateRequest")
	proto.RegisterType((*ModelCreateReply)(nil), "api.pb.ModelCreateReply")
	proto.RegisterType((*ListenRequest)(nil), "api.pb.ListenRequest")
	proto.RegisterType((*ListenReply)(nil), "api.pb.ListenReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	NewStore(ctx context.Context, in *NewStoreRequest, opts ...grpc.CallOption) (*NewStoreReply, error)
	RegisterSchema(ctx context.Context, in *RegisterSchemaRequest, opts ...grpc.CallOption) (*RegisterSchemaReply, error)
	ModelCreate(ctx context.Context, in *ModelCreateRequest, opts ...grpc.CallOption) (*ModelCreateReply, error)
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (API_ListenClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) NewStore(ctx context.Context, in *NewStoreRequest, opts ...grpc.CallOption) (*NewStoreReply, error) {
	out := new(NewStoreReply)
	err := c.cc.Invoke(ctx, "/api.pb.API/NewStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) RegisterSchema(ctx context.Context, in *RegisterSchemaRequest, opts ...grpc.CallOption) (*RegisterSchemaReply, error) {
	out := new(RegisterSchemaReply)
	err := c.cc.Invoke(ctx, "/api.pb.API/RegisterSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ModelCreate(ctx context.Context, in *ModelCreateRequest, opts ...grpc.CallOption) (*ModelCreateReply, error) {
	out := new(ModelCreateReply)
	err := c.cc.Invoke(ctx, "/api.pb.API/ModelCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (API_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/api.pb.API/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_ListenClient interface {
	Recv() (*ListenReply, error)
	grpc.ClientStream
}

type aPIListenClient struct {
	grpc.ClientStream
}

func (x *aPIListenClient) Recv() (*ListenReply, error) {
	m := new(ListenReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	NewStore(context.Context, *NewStoreRequest) (*NewStoreReply, error)
	RegisterSchema(context.Context, *RegisterSchemaRequest) (*RegisterSchemaReply, error)
	ModelCreate(context.Context, *ModelCreateRequest) (*ModelCreateReply, error)
	Listen(*ListenRequest, API_ListenServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_NewStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).NewStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pb.API/NewStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).NewStore(ctx, req.(*NewStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_RegisterSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).RegisterSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pb.API/RegisterSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).RegisterSchema(ctx, req.(*RegisterSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ModelCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ModelCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pb.API/ModelCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ModelCreate(ctx, req.(*ModelCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Listen(m, &aPIListenServer{stream})
}

type API_ListenServer interface {
	Send(*ListenReply) error
	grpc.ServerStream
}

type aPIListenServer struct {
	grpc.ServerStream
}

func (x *aPIListenServer) Send(m *ListenReply) error {
	return x.ServerStream.SendMsg(m)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.pb.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewStore",
			Handler:    _API_NewStore_Handler,
		},
		{
			MethodName: "RegisterSchema",
			Handler:    _API_RegisterSchema_Handler,
		},
		{
			MethodName: "ModelCreate",
			Handler:    _API_ModelCreate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _API_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_120f3cfc08b047b0) }

var fileDescriptor_api_120f3cfc08b047b0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xb5, 0xad, 0xa9, 0x30, 0x04, 0xd4, 0x21, 0xe0, 0xa6, 0x6a, 0x24, 0x9b, 0x98, 0x78, 0x6a,
	0x8c, 0x5e, 0x3c, 0x78, 0x31, 0xe2, 0xa1, 0x89, 0x12, 0x03, 0x67, 0x0f, 0x05, 0x26, 0xda, 0xa4,
	0xd0, 0xda, 0x2e, 0x1a, 0xfe, 0xc3, 0x0f, 0x36, 0xbb, 0xdb, 0x2d, 0x52, 0x30, 0x1e, 0xbc, 0xed,
	0x9b, 0x79, 0x33, 0x6f, 0xfa, 0x5e, 0xa1, 0x1e, 0xa6, 0x91, 0x9f, 0x66, 0x89, 0x48, 0xd0, 0x55,
	0xcf, 0x31, 0x3f, 0x84, 0xfd, 0x01, 0x7d, 0x8e, 0x44, 0x92, 0xd1, 0x90, 0xde, 0x17, 0x94, 0x0b,
	0x7e, 0x06, 0xcd, 0x55, 0x29, 0x8d, 0x97, 0xd8, 0x02, 0x3b, 0xe8, 0x33, 0xab, 0x67, 0x5d, 0xd4,
	0x87, 0x76, 0xd0, 0xe7, 0x2f, 0xd0, 0x19, 0xd2, 0x6b, 0x94, 0x0b, 0xca, 0x46, 0x93, 0x37, 0x9a,
	0x85, 0xc5, 0x24, 0x32, 0xd8, 0xcb, 0xe5, 0x58, 0xc9, 0x36, 0x10, 0x11, 0x76, 0xe7, 0xe1, 0x8c,
	0x98, 0xad, 0xca, 0xea, 0x8d, 0x5d, 0x70, 0x73, 0x35, 0xce, 0x1c, 0x55, 0x2d, 0x10, 0xef, 0x40,
	0xbb, 0xba, 0x3e, 0x8d, 0x97, 0x7c, 0x0a, 0xf8, 0x94, 0x4c, 0x29, 0xbe, 0xcf, 0x28, 0x14, 0xf4,
	0xb7, 0xe4, 0x09, 0xd4, 0x67, 0x92, 0x3f, 0x58, 0xe9, 0xae, 0x0a, 0x52, 0xfc, 0x23, 0x8c, 0x17,
	0x94, 0x33, 0xa7, 0xe7, 0x48, 0x71, 0x8d, 0xb8, 0x0f, 0x07, 0x6b, 0x2a, 0xf2, 0xfb, 0x3d, 0xa8,
	0xd1, 0x5c, 0x44, 0x22, 0xa2, 0x9c, 0x59, 0x8a, 0x5d, 0x62, 0x3e, 0x81, 0xe6, 0xa3, 0x3c, 0x75,
	0xfe, 0xdf, 0x83, 0x8c, 0xc8, 0x32, 0xe8, 0x17, 0x7e, 0x94, 0x98, 0x9f, 0x43, 0xc3, 0x88, 0xc8,
	0x7b, 0xba, 0xe0, 0xea, 0x56, 0xa1, 0x50, 0xa0, 0xab, 0x2f, 0x1b, 0x9c, 0xbb, 0xe7, 0x00, 0x6f,
	0xa1, 0x66, 0x02, 0xc4, 0x23, 0x5f, 0x07, 0xed, 0x57, 0x52, 0xf6, 0x3a, 0x9b, 0x0d, 0xe9, 0xf2,
	0x0e, 0x0e, 0xa0, 0xb5, 0x6e, 0x3f, 0x9e, 0x1a, 0xea, 0xd6, 0xd4, 0xbd, 0xe3, 0xdf, 0xda, 0x7a,
	0xdf, 0x03, 0x34, 0x7e, 0x38, 0x8a, 0x9e, 0x61, 0x6f, 0x86, 0xe9, 0xb1, 0xad, 0x3d, 0xbd, 0xe6,
	0x06, 0x5c, 0xed, 0x01, 0x96, 0x97, 0xaf, 0x19, 0xef, 0xb5, 0xab, 0x65, 0x35, 0x77, 0x69, 0x8d,
	0x5d, 0xf5, 0xc7, 0x5f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x26, 0x3f, 0x8f, 0xc1, 0xfe, 0x02,
	0x00, 0x00,
}
