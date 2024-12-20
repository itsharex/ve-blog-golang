// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: syslog.proto

// proto 包名

package syslogrpc

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
	SyslogRpc_AddOperationLog_FullMethodName      = "/syslogrpc.SyslogRpc/AddOperationLog"
	SyslogRpc_DeleteOperationLog_FullMethodName   = "/syslogrpc.SyslogRpc/DeleteOperationLog"
	SyslogRpc_FindOperationLogList_FullMethodName = "/syslogrpc.SyslogRpc/FindOperationLogList"
)

// SyslogRpcClient is the client API for SyslogRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyslogRpcClient interface {
	// 创建操作记录
	AddOperationLog(ctx context.Context, in *OperationLogNewReq, opts ...grpc.CallOption) (*OperationLogDetails, error)
	// 批量删除操作记录
	DeleteOperationLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询操作记录列表
	FindOperationLogList(ctx context.Context, in *FindOperationLogListReq, opts ...grpc.CallOption) (*FindOperationLogListResp, error)
}

type syslogRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewSyslogRpcClient(cc grpc.ClientConnInterface) SyslogRpcClient {
	return &syslogRpcClient{cc}
}

func (c *syslogRpcClient) AddOperationLog(ctx context.Context, in *OperationLogNewReq, opts ...grpc.CallOption) (*OperationLogDetails, error) {
	out := new(OperationLogDetails)
	err := c.cc.Invoke(ctx, SyslogRpc_AddOperationLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) DeleteOperationLog(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, SyslogRpc_DeleteOperationLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syslogRpcClient) FindOperationLogList(ctx context.Context, in *FindOperationLogListReq, opts ...grpc.CallOption) (*FindOperationLogListResp, error) {
	out := new(FindOperationLogListResp)
	err := c.cc.Invoke(ctx, SyslogRpc_FindOperationLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyslogRpcServer is the server API for SyslogRpc service.
// All implementations must embed UnimplementedSyslogRpcServer
// for forward compatibility
type SyslogRpcServer interface {
	// 创建操作记录
	AddOperationLog(context.Context, *OperationLogNewReq) (*OperationLogDetails, error)
	// 批量删除操作记录
	DeleteOperationLog(context.Context, *IdsReq) (*BatchResp, error)
	// 查询操作记录列表
	FindOperationLogList(context.Context, *FindOperationLogListReq) (*FindOperationLogListResp, error)
	mustEmbedUnimplementedSyslogRpcServer()
}

// UnimplementedSyslogRpcServer must be embedded to have forward compatible implementations.
type UnimplementedSyslogRpcServer struct {
}

func (UnimplementedSyslogRpcServer) AddOperationLog(context.Context, *OperationLogNewReq) (*OperationLogDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOperationLog not implemented")
}
func (UnimplementedSyslogRpcServer) DeleteOperationLog(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOperationLog not implemented")
}
func (UnimplementedSyslogRpcServer) FindOperationLogList(context.Context, *FindOperationLogListReq) (*FindOperationLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOperationLogList not implemented")
}
func (UnimplementedSyslogRpcServer) mustEmbedUnimplementedSyslogRpcServer() {}

// UnsafeSyslogRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyslogRpcServer will
// result in compilation errors.
type UnsafeSyslogRpcServer interface {
	mustEmbedUnimplementedSyslogRpcServer()
}

func RegisterSyslogRpcServer(s grpc.ServiceRegistrar, srv SyslogRpcServer) {
	s.RegisterService(&SyslogRpc_ServiceDesc, srv)
}

func _SyslogRpc_AddOperationLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationLogNewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).AddOperationLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_AddOperationLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).AddOperationLog(ctx, req.(*OperationLogNewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_DeleteOperationLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).DeleteOperationLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_DeleteOperationLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).DeleteOperationLog(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyslogRpc_FindOperationLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOperationLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyslogRpcServer).FindOperationLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyslogRpc_FindOperationLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyslogRpcServer).FindOperationLogList(ctx, req.(*FindOperationLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SyslogRpc_ServiceDesc is the grpc.ServiceDesc for SyslogRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyslogRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "syslogrpc.SyslogRpc",
	HandlerType: (*SyslogRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOperationLog",
			Handler:    _SyslogRpc_AddOperationLog_Handler,
		},
		{
			MethodName: "DeleteOperationLog",
			Handler:    _SyslogRpc_DeleteOperationLog_Handler,
		},
		{
			MethodName: "FindOperationLogList",
			Handler:    _SyslogRpc_FindOperationLogList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syslog.proto",
}
