// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: todo.proto

package todo_grpc_go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TodoService_CreateTodo_FullMethodName  = "/TodoService/CreateTodo"
	TodoService_GetTodoList_FullMethodName = "/TodoService/GetTodoList"
	TodoService_DeleteTodo_FullMethodName  = "/TodoService/DeleteTodo"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	GetTodoList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TodoService_GetTodoListClient, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, TodoService_CreateTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodoList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TodoService_GetTodoListClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoService_ServiceDesc.Streams[0], TodoService_GetTodoList_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceGetTodoListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoService_GetTodoListClient interface {
	Recv() (*Todo, error)
	grpc.ClientStream
}

type todoServiceGetTodoListClient struct {
	grpc.ClientStream
}

func (x *todoServiceGetTodoListClient) Recv() (*Todo, error) {
	m := new(Todo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, TodoService_DeleteTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreateTodo(context.Context, *Todo) (*Todo, error)
	GetTodoList(*emptypb.Empty, TodoService_GetTodoListServer) error
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreateTodo(context.Context, *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoServiceServer) GetTodoList(*emptypb.Empty, TodoService_GetTodoListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTodoList not implemented")
}
func (UnimplementedTodoServiceServer) DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_CreateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodoList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServiceServer).GetTodoList(m, &todoServiceGetTodoListServer{stream})
}

type TodoService_GetTodoListServer interface {
	Send(*Todo) error
	grpc.ServerStream
}

type todoServiceGetTodoListServer struct {
	grpc.ServerStream
}

func (x *todoServiceGetTodoListServer) Send(m *Todo) error {
	return x.ServerStream.SendMsg(m)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_DeleteTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTodoList",
			Handler:       _TodoService_GetTodoList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}
