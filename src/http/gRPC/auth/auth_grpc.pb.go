// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// Validate password, create and return a valid Token
	Login(ctx context.Context, in *LoginInput, opts ...grpc.CallOption) (*AuthOutput, error)
	// Validate a token and return a new one if the current token is valid
	Refresh(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*AuthOutput, error)
	// Make some token invalid (expire sessions)
	Invalidate(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*empty.Empty, error)
	// Invalidate all sessions for a given user
	InvalidateAll(ctx context.Context, in *SessionsInput, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get all user active sessions
	ActiveSessions(ctx context.Context, in *SessionsInput, opts ...grpc.CallOption) (*SessionsOutput, error)
	// Check if some token is valid or not
	IsValid(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginInput, opts ...grpc.CallOption) (*AuthOutput, error) {
	out := new(AuthOutput)
	err := c.cc.Invoke(ctx, "/AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Refresh(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*AuthOutput, error) {
	out := new(AuthOutput)
	err := c.cc.Invoke(ctx, "/AuthService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Invalidate(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/Invalidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) InvalidateAll(ctx context.Context, in *SessionsInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/InvalidateAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ActiveSessions(ctx context.Context, in *SessionsInput, opts ...grpc.CallOption) (*SessionsOutput, error) {
	out := new(SessionsOutput)
	err := c.cc.Invoke(ctx, "/AuthService/ActiveSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsValid(ctx context.Context, in *TokenInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/IsValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// Validate password, create and return a valid Token
	Login(context.Context, *LoginInput) (*AuthOutput, error)
	// Validate a token and return a new one if the current token is valid
	Refresh(context.Context, *TokenInput) (*AuthOutput, error)
	// Make some token invalid (expire sessions)
	Invalidate(context.Context, *TokenInput) (*empty.Empty, error)
	// Invalidate all sessions for a given user
	InvalidateAll(context.Context, *SessionsInput) (*empty.Empty, error)
	// Get all user active sessions
	ActiveSessions(context.Context, *SessionsInput) (*SessionsOutput, error)
	// Check if some token is valid or not
	IsValid(context.Context, *TokenInput) (*empty.Empty, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginInput) (*AuthOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Refresh(context.Context, *TokenInput) (*AuthOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthServiceServer) Invalidate(context.Context, *TokenInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invalidate not implemented")
}
func (UnimplementedAuthServiceServer) InvalidateAll(context.Context, *SessionsInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateAll not implemented")
}
func (UnimplementedAuthServiceServer) ActiveSessions(context.Context, *SessionsInput) (*SessionsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveSessions not implemented")
}
func (UnimplementedAuthServiceServer) IsValid(context.Context, *TokenInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValid not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Refresh(ctx, req.(*TokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Invalidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Invalidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/Invalidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Invalidate(ctx, req.(*TokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_InvalidateAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).InvalidateAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/InvalidateAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).InvalidateAll(ctx, req.(*SessionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ActiveSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ActiveSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/ActiveSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ActiveSessions(ctx, req.(*SessionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/IsValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsValid(ctx, req.(*TokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _AuthService_Refresh_Handler,
		},
		{
			MethodName: "Invalidate",
			Handler:    _AuthService_Invalidate_Handler,
		},
		{
			MethodName: "InvalidateAll",
			Handler:    _AuthService_InvalidateAll_Handler,
		},
		{
			MethodName: "ActiveSessions",
			Handler:    _AuthService_ActiveSessions_Handler,
		},
		{
			MethodName: "IsValid",
			Handler:    _AuthService_IsValid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/auth.proto",
}
