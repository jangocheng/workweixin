// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todos.proto

package todos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ToDoRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	ToDoID               uint64   `protobuf:"varint,3,opt,name=ToDoID,proto3" json:"ToDoID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToDoRequest) Reset()         { *m = ToDoRequest{} }
func (m *ToDoRequest) String() string { return proto.CompactTextString(m) }
func (*ToDoRequest) ProtoMessage()    {}
func (*ToDoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{0}
}

func (m *ToDoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToDoRequest.Unmarshal(m, b)
}
func (m *ToDoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToDoRequest.Marshal(b, m, deterministic)
}
func (m *ToDoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToDoRequest.Merge(m, src)
}
func (m *ToDoRequest) XXX_Size() int {
	return xxx_messageInfo_ToDoRequest.Size(m)
}
func (m *ToDoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ToDoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ToDoRequest proto.InternalMessageInfo

func (m *ToDoRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ToDoRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ToDoRequest) GetToDoID() uint64 {
	if m != nil {
		return m.ToDoID
	}
	return 0
}

type ToDoResult struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CreateTime           string   `protobuf:"bytes,3,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	FinishTime           string   `protobuf:"bytes,4,opt,name=FinishTime,proto3" json:"FinishTime,omitempty"`
	Active               bool     `protobuf:"varint,5,opt,name=Active,proto3" json:"Active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToDoResult) Reset()         { *m = ToDoResult{} }
func (m *ToDoResult) String() string { return proto.CompactTextString(m) }
func (*ToDoResult) ProtoMessage()    {}
func (*ToDoResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{1}
}

func (m *ToDoResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToDoResult.Unmarshal(m, b)
}
func (m *ToDoResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToDoResult.Marshal(b, m, deterministic)
}
func (m *ToDoResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToDoResult.Merge(m, src)
}
func (m *ToDoResult) XXX_Size() int {
	return xxx_messageInfo_ToDoResult.Size(m)
}
func (m *ToDoResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ToDoResult.DiscardUnknown(m)
}

var xxx_messageInfo_ToDoResult proto.InternalMessageInfo

func (m *ToDoResult) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ToDoResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ToDoResult) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ToDoResult) GetFinishTime() string {
	if m != nil {
		return m.FinishTime
	}
	return ""
}

func (m *ToDoResult) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type ToDoResponse struct {
	Result               []*ToDoResult `protobuf:"bytes,1,rep,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ToDoResponse) Reset()         { *m = ToDoResponse{} }
func (m *ToDoResponse) String() string { return proto.CompactTextString(m) }
func (*ToDoResponse) ProtoMessage()    {}
func (*ToDoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{2}
}

func (m *ToDoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToDoResponse.Unmarshal(m, b)
}
func (m *ToDoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToDoResponse.Marshal(b, m, deterministic)
}
func (m *ToDoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToDoResponse.Merge(m, src)
}
func (m *ToDoResponse) XXX_Size() int {
	return xxx_messageInfo_ToDoResponse.Size(m)
}
func (m *ToDoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ToDoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ToDoResponse proto.InternalMessageInfo

func (m *ToDoResponse) GetResult() []*ToDoResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ToDoRequest)(nil), "todos.ToDoRequest")
	proto.RegisterType((*ToDoResult)(nil), "todos.ToDoResult")
	proto.RegisterType((*ToDoResponse)(nil), "todos.ToDoResponse")
}

func init() { proto.RegisterFile("todos.proto", fileDescriptor_a7703e18c08e8710) }

var fileDescriptor_a7703e18c08e8710 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0x34, 0x4d, 0xed, 0x44, 0x04, 0x47, 0x90, 0xc5, 0x83, 0x84, 0x9c, 0xe2, 0xa5,
	0x87, 0xf6, 0xe4, 0x45, 0x90, 0x04, 0x21, 0x17, 0x0f, 0x6b, 0xc5, 0x73, 0x6d, 0x07, 0x0c, 0x34,
	0xd9, 0x9a, 0x9d, 0xfa, 0x1b, 0xfc, 0xd9, 0xb2, 0x93, 0x95, 0x9a, 0x5b, 0xde, 0x7b, 0x93, 0x6f,
	0x1f, 0x0f, 0x52, 0xb6, 0x3b, 0xeb, 0x16, 0x87, 0xde, 0xb2, 0xc5, 0xa9, 0x88, 0xfc, 0x1d, 0xd2,
	0xb5, 0xad, 0xac, 0xa1, 0xaf, 0x23, 0x39, 0xc6, 0x1b, 0x48, 0xde, 0x1c, 0xf5, 0x75, 0xa5, 0x55,
	0xa6, 0x8a, 0xb9, 0x09, 0x0a, 0x35, 0xcc, 0x4a, 0xdb, 0x31, 0x75, 0xac, 0x23, 0x09, 0xfe, 0xa4,
	0xff, 0xc3, 0x03, 0xea, 0x4a, 0x4f, 0x32, 0x55, 0xc4, 0x26, 0xa8, 0xfc, 0x47, 0x01, 0x0c, 0x64,
	0x77, 0xdc, 0x33, 0x5e, 0x42, 0x14, 0xa0, 0xb1, 0x89, 0xea, 0x0a, 0x11, 0xe2, 0x97, 0x4d, 0x4b,
	0x81, 0x26, 0xdf, 0x78, 0x07, 0x50, 0xf6, 0xb4, 0x61, 0x5a, 0x37, 0x2d, 0x09, 0x6e, 0x6e, 0xfe,
	0x39, 0x3e, 0x7f, 0x6e, 0xba, 0xc6, 0x7d, 0x4a, 0x1e, 0x0f, 0xf9, 0xc9, 0xf1, 0x55, 0x9e, 0xb6,
	0xdc, 0x7c, 0x93, 0x9e, 0x66, 0xaa, 0x38, 0x37, 0x41, 0xe5, 0x0f, 0x70, 0x11, 0x9a, 0x1c, 0x6c,
	0xe7, 0x08, 0xef, 0x21, 0x19, 0x5a, 0x69, 0x95, 0x4d, 0x8a, 0x74, 0x79, 0xb5, 0x18, 0x86, 0x39,
	0xd5, 0x35, 0xe1, 0x60, 0xf9, 0x08, 0x33, 0xef, 0x96, 0xed, 0x0e, 0x57, 0x90, 0xbc, 0xd2, 0x9e,
	0xb6, 0x8c, 0x38, 0xba, 0x97, 0xe1, 0x6e, 0xaf, 0xc7, 0x0c, 0x79, 0x28, 0x3f, 0xfb, 0x48, 0x64,
	0xec, 0xd5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4e, 0x65, 0xb7, 0x7b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ToDoCmdClient is the client API for ToDoCmd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToDoCmdClient interface {
	Select(ctx context.Context, in *ToDoRequest, opts ...grpc.CallOption) (*ToDoResponse, error)
}

type toDoCmdClient struct {
	cc *grpc.ClientConn
}

func NewToDoCmdClient(cc *grpc.ClientConn) ToDoCmdClient {
	return &toDoCmdClient{cc}
}

func (c *toDoCmdClient) Select(ctx context.Context, in *ToDoRequest, opts ...grpc.CallOption) (*ToDoResponse, error) {
	out := new(ToDoResponse)
	err := c.cc.Invoke(ctx, "/todos.ToDoCmd/Select", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoCmdServer is the server API for ToDoCmd service.
type ToDoCmdServer interface {
	Select(context.Context, *ToDoRequest) (*ToDoResponse, error)
}

// UnimplementedToDoCmdServer can be embedded to have forward compatible implementations.
type UnimplementedToDoCmdServer struct {
}

func (*UnimplementedToDoCmdServer) Select(ctx context.Context, req *ToDoRequest) (*ToDoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Select not implemented")
}

func RegisterToDoCmdServer(s *grpc.Server, srv ToDoCmdServer) {
	s.RegisterService(&_ToDoCmd_serviceDesc, srv)
}

func _ToDoCmd_Select_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoCmdServer).Select(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.ToDoCmd/Select",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoCmdServer).Select(ctx, req.(*ToDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToDoCmd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todos.ToDoCmd",
	HandlerType: (*ToDoCmdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Select",
			Handler:    _ToDoCmd_Select_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todos.proto",
}