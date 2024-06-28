// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: customer.proto

package user_service

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
	CustomerService_Create_FullMethodName     = "/user_service.CustomerService/Create"
	CustomerService_GetByID_FullMethodName    = "/user_service.CustomerService/GetByID"
	CustomerService_GetList_FullMethodName    = "/user_service.CustomerService/GetList"
	CustomerService_Update_FullMethodName     = "/user_service.CustomerService/Update"
	CustomerService_Delete_FullMethodName     = "/user_service.CustomerService/Delete"
	CustomerService_GetByGmail_FullMethodName = "/user_service.CustomerService/GetByGmail"
)

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	Create(ctx context.Context, in *CreateCustomer, opts ...grpc.CallOption) (*CustomerPrimaryKey, error)
	GetByID(ctx context.Context, in *CustomerPrimaryKey, opts ...grpc.CallOption) (*Customer, error)
	GetList(ctx context.Context, in *GetListCustomerRequest, opts ...grpc.CallOption) (*GetListCustomerResponse, error)
	Update(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerResponse, error)
	Delete(ctx context.Context, in *CustomerPrimaryKey, opts ...grpc.CallOption) (*Empty, error)
	GetByGmail(ctx context.Context, in *CustomerGmail, opts ...grpc.CallOption) (*CustomerPrimaryKey, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) Create(ctx context.Context, in *CreateCustomer, opts ...grpc.CallOption) (*CustomerPrimaryKey, error) {
	out := new(CustomerPrimaryKey)
	err := c.cc.Invoke(ctx, CustomerService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetByID(ctx context.Context, in *CustomerPrimaryKey, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, CustomerService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetList(ctx context.Context, in *GetListCustomerRequest, opts ...grpc.CallOption) (*GetListCustomerResponse, error) {
	out := new(GetListCustomerResponse)
	err := c.cc.Invoke(ctx, CustomerService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Update(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerResponse, error) {
	out := new(UpdateCustomerResponse)
	err := c.cc.Invoke(ctx, CustomerService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Delete(ctx context.Context, in *CustomerPrimaryKey, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, CustomerService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetByGmail(ctx context.Context, in *CustomerGmail, opts ...grpc.CallOption) (*CustomerPrimaryKey, error) {
	out := new(CustomerPrimaryKey)
	err := c.cc.Invoke(ctx, CustomerService_GetByGmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	Create(context.Context, *CreateCustomer) (*CustomerPrimaryKey, error)
	GetByID(context.Context, *CustomerPrimaryKey) (*Customer, error)
	GetList(context.Context, *GetListCustomerRequest) (*GetListCustomerResponse, error)
	Update(context.Context, *UpdateCustomerRequest) (*UpdateCustomerResponse, error)
	Delete(context.Context, *CustomerPrimaryKey) (*Empty, error)
	GetByGmail(context.Context, *CustomerGmail) (*CustomerPrimaryKey, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) Create(context.Context, *CreateCustomer) (*CustomerPrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCustomerServiceServer) GetByID(context.Context, *CustomerPrimaryKey) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedCustomerServiceServer) GetList(context.Context, *GetListCustomerRequest) (*GetListCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedCustomerServiceServer) Update(context.Context, *UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCustomerServiceServer) Delete(context.Context, *CustomerPrimaryKey) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCustomerServiceServer) GetByGmail(context.Context, *CustomerGmail) (*CustomerPrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByGmail not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Create(ctx, req.(*CreateCustomer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetByID(ctx, req.(*CustomerPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetList(ctx, req.(*GetListCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Update(ctx, req.(*UpdateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Delete(ctx, req.(*CustomerPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetByGmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerGmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetByGmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_GetByGmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetByGmail(ctx, req.(*CustomerGmail))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CustomerService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _CustomerService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _CustomerService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CustomerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CustomerService_Delete_Handler,
		},
		{
			MethodName: "GetByGmail",
			Handler:    _CustomerService_GetByGmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
