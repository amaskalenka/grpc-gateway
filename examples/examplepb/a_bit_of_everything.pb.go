// Code generated by protoc-gen-go.
// source: examples/examplepb/a_bit_of_everything.proto
// DO NOT EDIT!

/*
Package examplepb is a generated protocol buffer package.

It is generated from these files:
	examples/examplepb/a_bit_of_everything.proto
	examples/examplepb/echo_service.proto
	examples/examplepb/flow_combination.proto

It has these top-level messages:
	ABitOfEverything
	EmptyMessage
	IdMessage
	SimpleMessage
	EmptyProto
	NonEmptyProto
	UnaryProto
	NestedProto
	SingleNestedProto
*/
package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import google_api1 "google/api"
import gengo_grpc_gateway_examples_sub "github.com/gengo/grpc-gateway/examples/sub"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ABitOfEverything struct {
	Uuid         string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Nested       []*ABitOfEverything_Nested `protobuf:"bytes,2,rep,name=nested" json:"nested,omitempty"`
	FloatValue   float32                    `protobuf:"fixed32,3,opt,name=float_value" json:"float_value,omitempty"`
	DoubleValue  float64                    `protobuf:"fixed64,4,opt,name=double_value" json:"double_value,omitempty"`
	Int64Value   int64                      `protobuf:"varint,5,opt,name=int64_value" json:"int64_value,omitempty"`
	Uint64Value  uint64                     `protobuf:"varint,6,opt,name=uint64_value" json:"uint64_value,omitempty"`
	Int32Value   int32                      `protobuf:"varint,7,opt,name=int32_value" json:"int32_value,omitempty"`
	Fixed64Value uint64                     `protobuf:"fixed64,8,opt,name=fixed64_value" json:"fixed64_value,omitempty"`
	Fixed32Value uint32                     `protobuf:"fixed32,9,opt,name=fixed32_value" json:"fixed32_value,omitempty"`
	BoolValue    bool                       `protobuf:"varint,10,opt,name=bool_value" json:"bool_value,omitempty"`
	StringValue  string                     `protobuf:"bytes,11,opt,name=string_value" json:"string_value,omitempty"`
	// TODO(yugui) add bytes_value
	Uint32Value uint32 `protobuf:"varint,13,opt,name=uint32_value" json:"uint32_value,omitempty"`
	// TODO(yugui) add enum_value
	Sfixed32Value int32 `protobuf:"fixed32,15,opt,name=sfixed32_value" json:"sfixed32_value,omitempty"`
	Sfixed64Value int64 `protobuf:"fixed64,16,opt,name=sfixed64_value" json:"sfixed64_value,omitempty"`
	Sint32Value   int32 `protobuf:"zigzag32,17,opt,name=sint32_value" json:"sint32_value,omitempty"`
	Sint64Value   int64 `protobuf:"zigzag64,18,opt,name=sint64_value" json:"sint64_value,omitempty"`
}

func (m *ABitOfEverything) Reset()         { *m = ABitOfEverything{} }
func (m *ABitOfEverything) String() string { return proto.CompactTextString(m) }
func (*ABitOfEverything) ProtoMessage()    {}

func (m *ABitOfEverything) GetNested() []*ABitOfEverything_Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

type ABitOfEverything_Nested struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Amount uint32 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *ABitOfEverything_Nested) Reset()         { *m = ABitOfEverything_Nested{} }
func (m *ABitOfEverything_Nested) String() string { return proto.CompactTextString(m) }
func (*ABitOfEverything_Nested) ProtoMessage()    {}

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()         { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()    {}

type IdMessage struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *IdMessage) Reset()         { *m = IdMessage{} }
func (m *IdMessage) String() string { return proto.CompactTextString(m) }
func (*IdMessage) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for ABitOfEverythingService service

type ABitOfEverythingServiceClient interface {
	Create(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	CreateBody(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	BulkCreate(ctx context.Context, opts ...grpc.CallOption) (ABitOfEverythingService_BulkCreateClient, error)
	Lookup(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*ABitOfEverything, error)
	List(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (ABitOfEverythingService_ListClient, error)
	Update(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*EmptyMessage, error)
	Delete(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
	Echo(ctx context.Context, in *gengo_grpc_gateway_examples_sub.StringMessage, opts ...grpc.CallOption) (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	BulkEcho(ctx context.Context, opts ...grpc.CallOption) (ABitOfEverythingService_BulkEchoClient, error)
}

type aBitOfEverythingServiceClient struct {
	cc *grpc.ClientConn
}

func NewABitOfEverythingServiceClient(cc *grpc.ClientConn) ABitOfEverythingServiceClient {
	return &aBitOfEverythingServiceClient{cc}
}

func (c *aBitOfEverythingServiceClient) Create(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) CreateBody(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/CreateBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) BulkCreate(ctx context.Context, opts ...grpc.CallOption) (ABitOfEverythingService_BulkCreateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ABitOfEverythingService_serviceDesc.Streams[0], c.cc, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/BulkCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &aBitOfEverythingServiceBulkCreateClient{stream}
	return x, nil
}

type ABitOfEverythingService_BulkCreateClient interface {
	Send(*ABitOfEverything) error
	CloseAndRecv() (*EmptyMessage, error)
	grpc.ClientStream
}

type aBitOfEverythingServiceBulkCreateClient struct {
	grpc.ClientStream
}

func (x *aBitOfEverythingServiceBulkCreateClient) Send(m *ABitOfEverything) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aBitOfEverythingServiceBulkCreateClient) CloseAndRecv() (*EmptyMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aBitOfEverythingServiceClient) Lookup(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) List(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (ABitOfEverythingService_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ABitOfEverythingService_serviceDesc.Streams[1], c.cc, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &aBitOfEverythingServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ABitOfEverythingService_ListClient interface {
	Recv() (*ABitOfEverything, error)
	grpc.ClientStream
}

type aBitOfEverythingServiceListClient struct {
	grpc.ClientStream
}

func (x *aBitOfEverythingServiceListClient) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aBitOfEverythingServiceClient) Update(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Delete(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Echo(ctx context.Context, in *gengo_grpc_gateway_examples_sub.StringMessage, opts ...grpc.CallOption) (*gengo_grpc_gateway_examples_sub.StringMessage, error) {
	out := new(gengo_grpc_gateway_examples_sub.StringMessage)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) BulkEcho(ctx context.Context, opts ...grpc.CallOption) (ABitOfEverythingService_BulkEchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ABitOfEverythingService_serviceDesc.Streams[2], c.cc, "/gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService/BulkEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &aBitOfEverythingServiceBulkEchoClient{stream}
	return x, nil
}

type ABitOfEverythingService_BulkEchoClient interface {
	Send(*gengo_grpc_gateway_examples_sub.StringMessage) error
	Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	grpc.ClientStream
}

type aBitOfEverythingServiceBulkEchoClient struct {
	grpc.ClientStream
}

func (x *aBitOfEverythingServiceBulkEchoClient) Send(m *gengo_grpc_gateway_examples_sub.StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aBitOfEverythingServiceBulkEchoClient) Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error) {
	m := new(gengo_grpc_gateway_examples_sub.StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ABitOfEverythingService service

type ABitOfEverythingServiceServer interface {
	Create(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CreateBody(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	BulkCreate(ABitOfEverythingService_BulkCreateServer) error
	Lookup(context.Context, *IdMessage) (*ABitOfEverything, error)
	List(*EmptyMessage, ABitOfEverythingService_ListServer) error
	Update(context.Context, *ABitOfEverything) (*EmptyMessage, error)
	Delete(context.Context, *IdMessage) (*EmptyMessage, error)
	Echo(context.Context, *gengo_grpc_gateway_examples_sub.StringMessage) (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	BulkEcho(ABitOfEverythingService_BulkEchoServer) error
}

func RegisterABitOfEverythingServiceServer(s *grpc.Server, srv ABitOfEverythingServiceServer) {
	s.RegisterService(&_ABitOfEverythingService_serviceDesc, srv)
}

func _ABitOfEverythingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ABitOfEverythingServiceServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ABitOfEverythingService_CreateBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ABitOfEverythingServiceServer).CreateBody(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ABitOfEverythingService_BulkCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ABitOfEverythingServiceServer).BulkCreate(&aBitOfEverythingServiceBulkCreateServer{stream})
}

type ABitOfEverythingService_BulkCreateServer interface {
	SendAndClose(*EmptyMessage) error
	Recv() (*ABitOfEverything, error)
	grpc.ServerStream
}

type aBitOfEverythingServiceBulkCreateServer struct {
	grpc.ServerStream
}

func (x *aBitOfEverythingServiceBulkCreateServer) SendAndClose(m *EmptyMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aBitOfEverythingServiceBulkCreateServer) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ABitOfEverythingService_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ABitOfEverythingServiceServer).Lookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ABitOfEverythingService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ABitOfEverythingServiceServer).List(m, &aBitOfEverythingServiceListServer{stream})
}

type ABitOfEverythingService_ListServer interface {
	Send(*ABitOfEverything) error
	grpc.ServerStream
}

type aBitOfEverythingServiceListServer struct {
	grpc.ServerStream
}

func (x *aBitOfEverythingServiceListServer) Send(m *ABitOfEverything) error {
	return x.ServerStream.SendMsg(m)
}

func _ABitOfEverythingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ABitOfEverythingServiceServer).Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ABitOfEverythingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ABitOfEverythingServiceServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ABitOfEverythingService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(gengo_grpc_gateway_examples_sub.StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ABitOfEverythingServiceServer).Echo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ABitOfEverythingService_BulkEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ABitOfEverythingServiceServer).BulkEcho(&aBitOfEverythingServiceBulkEchoServer{stream})
}

type ABitOfEverythingService_BulkEchoServer interface {
	Send(*gengo_grpc_gateway_examples_sub.StringMessage) error
	Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	grpc.ServerStream
}

type aBitOfEverythingServiceBulkEchoServer struct {
	grpc.ServerStream
}

func (x *aBitOfEverythingServiceBulkEchoServer) Send(m *gengo_grpc_gateway_examples_sub.StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aBitOfEverythingServiceBulkEchoServer) Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error) {
	m := new(gengo_grpc_gateway_examples_sub.StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ABitOfEverythingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gengo.grpc.gateway.examples.examplepb.ABitOfEverythingService",
	HandlerType: (*ABitOfEverythingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ABitOfEverythingService_Create_Handler,
		},
		{
			MethodName: "CreateBody",
			Handler:    _ABitOfEverythingService_CreateBody_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _ABitOfEverythingService_Lookup_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ABitOfEverythingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ABitOfEverythingService_Delete_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _ABitOfEverythingService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreate",
			Handler:       _ABitOfEverythingService_BulkCreate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _ABitOfEverythingService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BulkEcho",
			Handler:       _ABitOfEverythingService_BulkEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
