// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: system_user.proto

package auth_service

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
	SystemUserAuth_SystemUserLoginByPassword_FullMethodName       = "/auth_service.SystemUserAuth/SystemUserLoginByPassword"
	SystemUserAuth_SystemUserGmailCheck_FullMethodName            = "/auth_service.SystemUserAuth/SystemUserGmailCheck"
	SystemUserAuth_SystemUserRegisterByMail_FullMethodName        = "/auth_service.SystemUserAuth/SystemUserRegisterByMail"
	SystemUserAuth_SystemUserRegisterByMailConfirm_FullMethodName = "/auth_service.SystemUserAuth/SystemUserRegisterByMailConfirm"
	SystemUserAuth_SystemUserCreate_FullMethodName                = "/auth_service.SystemUserAuth/SystemUserCreate"
	SystemUserAuth_SystemUserLoginByGmail_FullMethodName          = "/auth_service.SystemUserAuth/SystemUserLoginByGmail"
	SystemUserAuth_SystemUserLoginByGmailComfirm_FullMethodName   = "/auth_service.SystemUserAuth/SystemUserLoginByGmailComfirm"
	SystemUserAuth_SystemUserUpdatePassword_FullMethodName        = "/auth_service.SystemUserAuth/SystemUserUpdatePassword"
	SystemUserAuth_SystemUserResetPassword_FullMethodName         = "/auth_service.SystemUserAuth/SystemUserResetPassword"
	SystemUserAuth_SystemUserResetPasswordConfirm_FullMethodName  = "/auth_service.SystemUserAuth/SystemUserResetPasswordConfirm"
)

// SystemUserAuthClient is the client API for SystemUserAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemUserAuthClient interface {
	SystemUserLoginByPassword(ctx context.Context, in *SystemUserLoginRequest, opts ...grpc.CallOption) (*SystemUserLoginResponse, error)
	SystemUserGmailCheck(ctx context.Context, in *SystemUserGmailCheckRequest, opts ...grpc.CallOption) (*SystemUserGmailCheckResponse, error)
	SystemUserRegisterByMail(ctx context.Context, in *SystemUserGmailCheckRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error)
	SystemUserRegisterByMailConfirm(ctx context.Context, in *SystemUserRConfirm, opts ...grpc.CallOption) (*RespRegSeller, error)
	SystemUserCreate(ctx context.Context, in *SystemUserCreateRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error)
	SystemUserLoginByGmail(ctx context.Context, in *SystemUserGmailCheckRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error)
	SystemUserLoginByGmailComfirm(ctx context.Context, in *SystemUserLoginByGmailRequest, opts ...grpc.CallOption) (*SystemUserLoginResponse, error)
	SystemUserUpdatePassword(ctx context.Context, in *SystemUserCreateRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error)
	SystemUserResetPassword(ctx context.Context, in *SystemUserGmailCheckRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error)
	SystemUserResetPasswordConfirm(ctx context.Context, in *SystemUserPasswordConfirm, opts ...grpc.CallOption) (*SystemUserEmpty, error)
}

type systemUserAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemUserAuthClient(cc grpc.ClientConnInterface) SystemUserAuthClient {
	return &systemUserAuthClient{cc}
}

func (c *systemUserAuthClient) SystemUserLoginByPassword(ctx context.Context, in *SystemUserLoginRequest, opts ...grpc.CallOption) (*SystemUserLoginResponse, error) {
	out := new(SystemUserLoginResponse)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserLoginByPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserGmailCheck(ctx context.Context, in *SystemUserGmailCheckRequest, opts ...grpc.CallOption) (*SystemUserGmailCheckResponse, error) {
	out := new(SystemUserGmailCheckResponse)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserGmailCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserRegisterByMail(ctx context.Context, in *SystemUserGmailCheckRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error) {
	out := new(SystemUserEmpty)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserRegisterByMail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserRegisterByMailConfirm(ctx context.Context, in *SystemUserRConfirm, opts ...grpc.CallOption) (*RespRegSeller, error) {
	out := new(RespRegSeller)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserRegisterByMailConfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserCreate(ctx context.Context, in *SystemUserCreateRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error) {
	out := new(SystemUserEmpty)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserLoginByGmail(ctx context.Context, in *SystemUserGmailCheckRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error) {
	out := new(SystemUserEmpty)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserLoginByGmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserLoginByGmailComfirm(ctx context.Context, in *SystemUserLoginByGmailRequest, opts ...grpc.CallOption) (*SystemUserLoginResponse, error) {
	out := new(SystemUserLoginResponse)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserLoginByGmailComfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserUpdatePassword(ctx context.Context, in *SystemUserCreateRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error) {
	out := new(SystemUserEmpty)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserUpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserResetPassword(ctx context.Context, in *SystemUserGmailCheckRequest, opts ...grpc.CallOption) (*SystemUserEmpty, error) {
	out := new(SystemUserEmpty)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserResetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserAuthClient) SystemUserResetPasswordConfirm(ctx context.Context, in *SystemUserPasswordConfirm, opts ...grpc.CallOption) (*SystemUserEmpty, error) {
	out := new(SystemUserEmpty)
	err := c.cc.Invoke(ctx, SystemUserAuth_SystemUserResetPasswordConfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemUserAuthServer is the server API for SystemUserAuth service.
// All implementations must embed UnimplementedSystemUserAuthServer
// for forward compatibility
type SystemUserAuthServer interface {
	SystemUserLoginByPassword(context.Context, *SystemUserLoginRequest) (*SystemUserLoginResponse, error)
	SystemUserGmailCheck(context.Context, *SystemUserGmailCheckRequest) (*SystemUserGmailCheckResponse, error)
	SystemUserRegisterByMail(context.Context, *SystemUserGmailCheckRequest) (*SystemUserEmpty, error)
	SystemUserRegisterByMailConfirm(context.Context, *SystemUserRConfirm) (*RespRegSeller, error)
	SystemUserCreate(context.Context, *SystemUserCreateRequest) (*SystemUserEmpty, error)
	SystemUserLoginByGmail(context.Context, *SystemUserGmailCheckRequest) (*SystemUserEmpty, error)
	SystemUserLoginByGmailComfirm(context.Context, *SystemUserLoginByGmailRequest) (*SystemUserLoginResponse, error)
	SystemUserUpdatePassword(context.Context, *SystemUserCreateRequest) (*SystemUserEmpty, error)
	SystemUserResetPassword(context.Context, *SystemUserGmailCheckRequest) (*SystemUserEmpty, error)
	SystemUserResetPasswordConfirm(context.Context, *SystemUserPasswordConfirm) (*SystemUserEmpty, error)
	mustEmbedUnimplementedSystemUserAuthServer()
}

// UnimplementedSystemUserAuthServer must be embedded to have forward compatible implementations.
type UnimplementedSystemUserAuthServer struct {
}

func (UnimplementedSystemUserAuthServer) SystemUserLoginByPassword(context.Context, *SystemUserLoginRequest) (*SystemUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserLoginByPassword not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserGmailCheck(context.Context, *SystemUserGmailCheckRequest) (*SystemUserGmailCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserGmailCheck not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserRegisterByMail(context.Context, *SystemUserGmailCheckRequest) (*SystemUserEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserRegisterByMail not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserRegisterByMailConfirm(context.Context, *SystemUserRConfirm) (*RespRegSeller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserRegisterByMailConfirm not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserCreate(context.Context, *SystemUserCreateRequest) (*SystemUserEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserCreate not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserLoginByGmail(context.Context, *SystemUserGmailCheckRequest) (*SystemUserEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserLoginByGmail not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserLoginByGmailComfirm(context.Context, *SystemUserLoginByGmailRequest) (*SystemUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserLoginByGmailComfirm not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserUpdatePassword(context.Context, *SystemUserCreateRequest) (*SystemUserEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserUpdatePassword not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserResetPassword(context.Context, *SystemUserGmailCheckRequest) (*SystemUserEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserResetPassword not implemented")
}
func (UnimplementedSystemUserAuthServer) SystemUserResetPasswordConfirm(context.Context, *SystemUserPasswordConfirm) (*SystemUserEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUserResetPasswordConfirm not implemented")
}
func (UnimplementedSystemUserAuthServer) mustEmbedUnimplementedSystemUserAuthServer() {}

// UnsafeSystemUserAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemUserAuthServer will
// result in compilation errors.
type UnsafeSystemUserAuthServer interface {
	mustEmbedUnimplementedSystemUserAuthServer()
}

func RegisterSystemUserAuthServer(s grpc.ServiceRegistrar, srv SystemUserAuthServer) {
	s.RegisterService(&SystemUserAuth_ServiceDesc, srv)
}

func _SystemUserAuth_SystemUserLoginByPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserLoginByPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserLoginByPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserLoginByPassword(ctx, req.(*SystemUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserGmailCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserGmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserGmailCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserGmailCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserGmailCheck(ctx, req.(*SystemUserGmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserRegisterByMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserGmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserRegisterByMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserRegisterByMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserRegisterByMail(ctx, req.(*SystemUserGmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserRegisterByMailConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserRConfirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserRegisterByMailConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserRegisterByMailConfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserRegisterByMailConfirm(ctx, req.(*SystemUserRConfirm))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserCreate(ctx, req.(*SystemUserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserLoginByGmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserGmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserLoginByGmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserLoginByGmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserLoginByGmail(ctx, req.(*SystemUserGmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserLoginByGmailComfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserLoginByGmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserLoginByGmailComfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserLoginByGmailComfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserLoginByGmailComfirm(ctx, req.(*SystemUserLoginByGmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserUpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserUpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserUpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserUpdatePassword(ctx, req.(*SystemUserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserGmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserResetPassword(ctx, req.(*SystemUserGmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserAuth_SystemUserResetPasswordConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserPasswordConfirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserAuthServer).SystemUserResetPasswordConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemUserAuth_SystemUserResetPasswordConfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserAuthServer).SystemUserResetPasswordConfirm(ctx, req.(*SystemUserPasswordConfirm))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemUserAuth_ServiceDesc is the grpc.ServiceDesc for SystemUserAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemUserAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.SystemUserAuth",
	HandlerType: (*SystemUserAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemUserLoginByPassword",
			Handler:    _SystemUserAuth_SystemUserLoginByPassword_Handler,
		},
		{
			MethodName: "SystemUserGmailCheck",
			Handler:    _SystemUserAuth_SystemUserGmailCheck_Handler,
		},
		{
			MethodName: "SystemUserRegisterByMail",
			Handler:    _SystemUserAuth_SystemUserRegisterByMail_Handler,
		},
		{
			MethodName: "SystemUserRegisterByMailConfirm",
			Handler:    _SystemUserAuth_SystemUserRegisterByMailConfirm_Handler,
		},
		{
			MethodName: "SystemUserCreate",
			Handler:    _SystemUserAuth_SystemUserCreate_Handler,
		},
		{
			MethodName: "SystemUserLoginByGmail",
			Handler:    _SystemUserAuth_SystemUserLoginByGmail_Handler,
		},
		{
			MethodName: "SystemUserLoginByGmailComfirm",
			Handler:    _SystemUserAuth_SystemUserLoginByGmailComfirm_Handler,
		},
		{
			MethodName: "SystemUserUpdatePassword",
			Handler:    _SystemUserAuth_SystemUserUpdatePassword_Handler,
		},
		{
			MethodName: "SystemUserResetPassword",
			Handler:    _SystemUserAuth_SystemUserResetPassword_Handler,
		},
		{
			MethodName: "SystemUserResetPasswordConfirm",
			Handler:    _SystemUserAuth_SystemUserResetPasswordConfirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_user.proto",
}