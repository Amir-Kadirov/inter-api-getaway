// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: attendance.proto

package schedule_service

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
	AttendanceService_Create_FullMethodName  = "/schedule_service.AttendanceService/Create"
	AttendanceService_GetByID_FullMethodName = "/schedule_service.AttendanceService/GetByID"
	AttendanceService_GetList_FullMethodName = "/schedule_service.AttendanceService/GetList"
)

// AttendanceServiceClient is the client API for AttendanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttendanceServiceClient interface {
	Create(ctx context.Context, in *CreateAttendance, opts ...grpc.CallOption) (*ATMessage, error)
	GetByID(ctx context.Context, in *AttendancePrimaryKey, opts ...grpc.CallOption) (*Attendance, error)
	GetList(ctx context.Context, in *GetListAttendanceRequest, opts ...grpc.CallOption) (*GetListAttendanceResponse, error)
}

type attendanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttendanceServiceClient(cc grpc.ClientConnInterface) AttendanceServiceClient {
	return &attendanceServiceClient{cc}
}

func (c *attendanceServiceClient) Create(ctx context.Context, in *CreateAttendance, opts ...grpc.CallOption) (*ATMessage, error) {
	out := new(ATMessage)
	err := c.cc.Invoke(ctx, AttendanceService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceServiceClient) GetByID(ctx context.Context, in *AttendancePrimaryKey, opts ...grpc.CallOption) (*Attendance, error) {
	out := new(Attendance)
	err := c.cc.Invoke(ctx, AttendanceService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceServiceClient) GetList(ctx context.Context, in *GetListAttendanceRequest, opts ...grpc.CallOption) (*GetListAttendanceResponse, error) {
	out := new(GetListAttendanceResponse)
	err := c.cc.Invoke(ctx, AttendanceService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttendanceServiceServer is the server API for AttendanceService service.
// All implementations must embed UnimplementedAttendanceServiceServer
// for forward compatibility
type AttendanceServiceServer interface {
	Create(context.Context, *CreateAttendance) (*ATMessage, error)
	GetByID(context.Context, *AttendancePrimaryKey) (*Attendance, error)
	GetList(context.Context, *GetListAttendanceRequest) (*GetListAttendanceResponse, error)
	mustEmbedUnimplementedAttendanceServiceServer()
}

// UnimplementedAttendanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAttendanceServiceServer struct {
}

func (UnimplementedAttendanceServiceServer) Create(context.Context, *CreateAttendance) (*ATMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAttendanceServiceServer) GetByID(context.Context, *AttendancePrimaryKey) (*Attendance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedAttendanceServiceServer) GetList(context.Context, *GetListAttendanceRequest) (*GetListAttendanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAttendanceServiceServer) mustEmbedUnimplementedAttendanceServiceServer() {}

// UnsafeAttendanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttendanceServiceServer will
// result in compilation errors.
type UnsafeAttendanceServiceServer interface {
	mustEmbedUnimplementedAttendanceServiceServer()
}

func RegisterAttendanceServiceServer(s grpc.ServiceRegistrar, srv AttendanceServiceServer) {
	s.RegisterService(&AttendanceService_ServiceDesc, srv)
}

func _AttendanceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttendance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).Create(ctx, req.(*CreateAttendance))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttendanceService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendancePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).GetByID(ctx, req.(*AttendancePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttendanceService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAttendanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).GetList(ctx, req.(*GetListAttendanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AttendanceService_ServiceDesc is the grpc.ServiceDesc for AttendanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttendanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule_service.AttendanceService",
	HandlerType: (*AttendanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AttendanceService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _AttendanceService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _AttendanceService_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attendance.proto",
}
