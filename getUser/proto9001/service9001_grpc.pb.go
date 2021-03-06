// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto9001

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

// UserService9001Client is the client API for UserService9001 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserService9001Client interface {
	GetUser(ctx context.Context, in *RequestFor9001, opts ...grpc.CallOption) (*ResponseFor9001, error)
}

type userService9001Client struct {
	cc grpc.ClientConnInterface
}

func NewUserService9001Client(cc grpc.ClientConnInterface) UserService9001Client {
	return &userService9001Client{cc}
}

func (c *userService9001Client) GetUser(ctx context.Context, in *RequestFor9001, opts ...grpc.CallOption) (*ResponseFor9001, error) {
	out := new(ResponseFor9001)
	err := c.cc.Invoke(ctx, "/proto9001.UserService9001/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserService9001Server is the server API for UserService9001 service.
// All implementations should embed UnimplementedUserService9001Server
// for forward compatibility
type UserService9001Server interface {
	GetUser(context.Context, *RequestFor9001) (*ResponseFor9001, error)
}

// UnimplementedUserService9001Server should be embedded to have forward compatible implementations.
type UnimplementedUserService9001Server struct {
}

func (UnimplementedUserService9001Server) GetUser(context.Context, *RequestFor9001) (*ResponseFor9001, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

// UnsafeUserService9001Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserService9001Server will
// result in compilation errors.
type UnsafeUserService9001Server interface {
	mustEmbedUnimplementedUserService9001Server()
}

func RegisterUserService9001Server(s grpc.ServiceRegistrar, srv UserService9001Server) {
	s.RegisterService(&UserService9001_ServiceDesc, srv)
}

func _UserService9001_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFor9001)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserService9001Server).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto9001.UserService9001/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserService9001Server).GetUser(ctx, req.(*RequestFor9001))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService9001_ServiceDesc is the grpc.ServiceDesc for UserService9001 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService9001_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto9001.UserService9001",
	HandlerType: (*UserService9001Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService9001_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service9001.proto",
}
