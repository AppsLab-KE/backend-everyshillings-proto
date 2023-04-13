// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: db/server.proto

package db

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DbService_HealthCheck_FullMethodName    = "/db.DbService/HealthCheck"
	DbService_CreateUser_FullMethodName     = "/db.DbService/CreateUser"
	DbService_UpdateUser_FullMethodName     = "/db.DbService/UpdateUser"
	DbService_GetPagedUsers_FullMethodName  = "/db.DbService/GetPagedUsers"
	DbService_GetUserByField_FullMethodName = "/db.DbService/GetUserByField"
)

// DbServiceClient is the client API for DbService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DbServiceClient interface {
	HealthCheck(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	// USERS
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRes, error)
	GetPagedUsers(ctx context.Context, in *GetPagedUsersReq, opts ...grpc.CallOption) (*GetPagedUsersRes, error)
	GetUserByField(ctx context.Context, in *GetByfieldReq, opts ...grpc.CallOption) (*GetByfieldRes, error)
}

type dbServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDbServiceClient(cc grpc.ClientConnInterface) DbServiceClient {
	return &dbServiceClient{cc}
}

func (c *dbServiceClient) HealthCheck(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, DbService_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error) {
	out := new(CreateUserRes)
	err := c.cc.Invoke(ctx, DbService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRes, error) {
	out := new(UpdateUserRes)
	err := c.cc.Invoke(ctx, DbService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) GetPagedUsers(ctx context.Context, in *GetPagedUsersReq, opts ...grpc.CallOption) (*GetPagedUsersRes, error) {
	out := new(GetPagedUsersRes)
	err := c.cc.Invoke(ctx, DbService_GetPagedUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbServiceClient) GetUserByField(ctx context.Context, in *GetByfieldReq, opts ...grpc.CallOption) (*GetByfieldRes, error) {
	out := new(GetByfieldRes)
	err := c.cc.Invoke(ctx, DbService_GetUserByField_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DbServiceServer is the server API for DbService service.
// All implementations must embed UnimplementedDbServiceServer
// for forward compatibility
type DbServiceServer interface {
	HealthCheck(context.Context, *DefaultRequest) (*HealthResponse, error)
	// USERS
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserRes, error)
	GetPagedUsers(context.Context, *GetPagedUsersReq) (*GetPagedUsersRes, error)
	GetUserByField(context.Context, *GetByfieldReq) (*GetByfieldRes, error)
	mustEmbedUnimplementedDbServiceServer()
}

// UnimplementedDbServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDbServiceServer struct {
}

func (UnimplementedDbServiceServer) HealthCheck(context.Context, *DefaultRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedDbServiceServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedDbServiceServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedDbServiceServer) GetPagedUsers(context.Context, *GetPagedUsersReq) (*GetPagedUsersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPagedUsers not implemented")
}
func (UnimplementedDbServiceServer) GetUserByField(context.Context, *GetByfieldReq) (*GetByfieldRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByField not implemented")
}
func (UnimplementedDbServiceServer) mustEmbedUnimplementedDbServiceServer() {}

// UnsafeDbServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DbServiceServer will
// result in compilation errors.
type UnsafeDbServiceServer interface {
	mustEmbedUnimplementedDbServiceServer()
}

func RegisterDbServiceServer(s grpc.ServiceRegistrar, srv DbServiceServer) {
	s.RegisterService(&DbService_ServiceDesc, srv)
}

func _DbService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbService_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).HealthCheck(ctx, req.(*DefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_GetPagedUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPagedUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).GetPagedUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbService_GetPagedUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).GetPagedUsers(ctx, req.(*GetPagedUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbService_GetUserByField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByfieldReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServiceServer).GetUserByField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbService_GetUserByField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServiceServer).GetUserByField(ctx, req.(*GetByfieldReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DbService_ServiceDesc is the grpc.ServiceDesc for DbService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DbService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.DbService",
	HandlerType: (*DbServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _DbService_HealthCheck_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _DbService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _DbService_UpdateUser_Handler,
		},
		{
			MethodName: "GetPagedUsers",
			Handler:    _DbService_GetPagedUsers_Handler,
		},
		{
			MethodName: "GetUserByField",
			Handler:    _DbService_GetUserByField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db/server.proto",
}
