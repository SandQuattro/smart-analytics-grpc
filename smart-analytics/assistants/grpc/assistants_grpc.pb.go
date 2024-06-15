// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/assistants/assistants.proto

package grpc

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
	AssistantService_ListAssistantFiles_FullMethodName  = "/assistants.AssistantService/ListAssistantFiles"
	AssistantService_DeleteAssistantFile_FullMethodName = "/assistants.AssistantService/DeleteAssistantFile"
	AssistantService_LinkFileToAssistant_FullMethodName = "/assistants.AssistantService/LinkFileToAssistant"
	AssistantService_CreateThread_FullMethodName        = "/assistants.AssistantService/CreateThread"
	AssistantService_GetThread_FullMethodName           = "/assistants.AssistantService/GetThread"
	AssistantService_DeleteThread_FullMethodName        = "/assistants.AssistantService/DeleteThread"
	AssistantService_CreateThreadRun_FullMethodName     = "/assistants.AssistantService/CreateThreadRun"
	AssistantService_GetThreadRuns_FullMethodName       = "/assistants.AssistantService/GetThreadRuns"
)

// AssistantServiceClient is the client API for AssistantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssistantServiceClient interface {
	ListAssistantFiles(ctx context.Context, in *AssistantRequest, opts ...grpc.CallOption) (*ListAssistantFilesResponse, error)
	DeleteAssistantFile(ctx context.Context, in *AssistantFileRequest, opts ...grpc.CallOption) (*DeletedObject, error)
	LinkFileToAssistant(ctx context.Context, in *AssistantFileRequest, opts ...grpc.CallOption) (*AssistantFileData, error)
	CreateThread(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AssistantObject, error)
	GetThread(ctx context.Context, in *ThreadRequest, opts ...grpc.CallOption) (*AssistantObject, error)
	DeleteThread(ctx context.Context, in *ThreadRequest, opts ...grpc.CallOption) (*DeletedObject, error)
	CreateThreadRun(ctx context.Context, in *CreateThreadRunRequest, opts ...grpc.CallOption) (*ThreadRun, error)
	GetThreadRuns(ctx context.Context, in *ThreadRequest, opts ...grpc.CallOption) (AssistantService_GetThreadRunsClient, error)
}

type assistantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssistantServiceClient(cc grpc.ClientConnInterface) AssistantServiceClient {
	return &assistantServiceClient{cc}
}

func (c *assistantServiceClient) ListAssistantFiles(ctx context.Context, in *AssistantRequest, opts ...grpc.CallOption) (*ListAssistantFilesResponse, error) {
	out := new(ListAssistantFilesResponse)
	err := c.cc.Invoke(ctx, AssistantService_ListAssistantFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistantServiceClient) DeleteAssistantFile(ctx context.Context, in *AssistantFileRequest, opts ...grpc.CallOption) (*DeletedObject, error) {
	out := new(DeletedObject)
	err := c.cc.Invoke(ctx, AssistantService_DeleteAssistantFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistantServiceClient) LinkFileToAssistant(ctx context.Context, in *AssistantFileRequest, opts ...grpc.CallOption) (*AssistantFileData, error) {
	out := new(AssistantFileData)
	err := c.cc.Invoke(ctx, AssistantService_LinkFileToAssistant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistantServiceClient) CreateThread(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AssistantObject, error) {
	out := new(AssistantObject)
	err := c.cc.Invoke(ctx, AssistantService_CreateThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistantServiceClient) GetThread(ctx context.Context, in *ThreadRequest, opts ...grpc.CallOption) (*AssistantObject, error) {
	out := new(AssistantObject)
	err := c.cc.Invoke(ctx, AssistantService_GetThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistantServiceClient) DeleteThread(ctx context.Context, in *ThreadRequest, opts ...grpc.CallOption) (*DeletedObject, error) {
	out := new(DeletedObject)
	err := c.cc.Invoke(ctx, AssistantService_DeleteThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistantServiceClient) CreateThreadRun(ctx context.Context, in *CreateThreadRunRequest, opts ...grpc.CallOption) (*ThreadRun, error) {
	out := new(ThreadRun)
	err := c.cc.Invoke(ctx, AssistantService_CreateThreadRun_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistantServiceClient) GetThreadRuns(ctx context.Context, in *ThreadRequest, opts ...grpc.CallOption) (AssistantService_GetThreadRunsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AssistantService_ServiceDesc.Streams[0], AssistantService_GetThreadRuns_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &assistantServiceGetThreadRunsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AssistantService_GetThreadRunsClient interface {
	Recv() (*ThreadRun, error)
	grpc.ClientStream
}

type assistantServiceGetThreadRunsClient struct {
	grpc.ClientStream
}

func (x *assistantServiceGetThreadRunsClient) Recv() (*ThreadRun, error) {
	m := new(ThreadRun)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AssistantServiceServer is the server API for AssistantService service.
// All implementations must embed UnimplementedAssistantServiceServer
// for forward compatibility
type AssistantServiceServer interface {
	ListAssistantFiles(context.Context, *AssistantRequest) (*ListAssistantFilesResponse, error)
	DeleteAssistantFile(context.Context, *AssistantFileRequest) (*DeletedObject, error)
	LinkFileToAssistant(context.Context, *AssistantFileRequest) (*AssistantFileData, error)
	CreateThread(context.Context, *emptypb.Empty) (*AssistantObject, error)
	GetThread(context.Context, *ThreadRequest) (*AssistantObject, error)
	DeleteThread(context.Context, *ThreadRequest) (*DeletedObject, error)
	CreateThreadRun(context.Context, *CreateThreadRunRequest) (*ThreadRun, error)
	GetThreadRuns(*ThreadRequest, AssistantService_GetThreadRunsServer) error
	mustEmbedUnimplementedAssistantServiceServer()
}

// UnimplementedAssistantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssistantServiceServer struct {
}

func (UnimplementedAssistantServiceServer) ListAssistantFiles(context.Context, *AssistantRequest) (*ListAssistantFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssistantFiles not implemented")
}
func (UnimplementedAssistantServiceServer) DeleteAssistantFile(context.Context, *AssistantFileRequest) (*DeletedObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssistantFile not implemented")
}
func (UnimplementedAssistantServiceServer) LinkFileToAssistant(context.Context, *AssistantFileRequest) (*AssistantFileData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkFileToAssistant not implemented")
}
func (UnimplementedAssistantServiceServer) CreateThread(context.Context, *emptypb.Empty) (*AssistantObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateThread not implemented")
}
func (UnimplementedAssistantServiceServer) GetThread(context.Context, *ThreadRequest) (*AssistantObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThread not implemented")
}
func (UnimplementedAssistantServiceServer) DeleteThread(context.Context, *ThreadRequest) (*DeletedObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteThread not implemented")
}
func (UnimplementedAssistantServiceServer) CreateThreadRun(context.Context, *CreateThreadRunRequest) (*ThreadRun, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateThreadRun not implemented")
}
func (UnimplementedAssistantServiceServer) GetThreadRuns(*ThreadRequest, AssistantService_GetThreadRunsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetThreadRuns not implemented")
}
func (UnimplementedAssistantServiceServer) mustEmbedUnimplementedAssistantServiceServer() {}

// UnsafeAssistantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssistantServiceServer will
// result in compilation errors.
type UnsafeAssistantServiceServer interface {
	mustEmbedUnimplementedAssistantServiceServer()
}

func RegisterAssistantServiceServer(s grpc.ServiceRegistrar, srv AssistantServiceServer) {
	s.RegisterService(&AssistantService_ServiceDesc, srv)
}

func _AssistantService_ListAssistantFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssistantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistantServiceServer).ListAssistantFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssistantService_ListAssistantFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistantServiceServer).ListAssistantFiles(ctx, req.(*AssistantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistantService_DeleteAssistantFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssistantFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistantServiceServer).DeleteAssistantFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssistantService_DeleteAssistantFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistantServiceServer).DeleteAssistantFile(ctx, req.(*AssistantFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistantService_LinkFileToAssistant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssistantFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistantServiceServer).LinkFileToAssistant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssistantService_LinkFileToAssistant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistantServiceServer).LinkFileToAssistant(ctx, req.(*AssistantFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistantService_CreateThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistantServiceServer).CreateThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssistantService_CreateThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistantServiceServer).CreateThread(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistantService_GetThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistantServiceServer).GetThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssistantService_GetThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistantServiceServer).GetThread(ctx, req.(*ThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistantService_DeleteThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistantServiceServer).DeleteThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssistantService_DeleteThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistantServiceServer).DeleteThread(ctx, req.(*ThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistantService_CreateThreadRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateThreadRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistantServiceServer).CreateThreadRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssistantService_CreateThreadRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistantServiceServer).CreateThreadRun(ctx, req.(*CreateThreadRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistantService_GetThreadRuns_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ThreadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AssistantServiceServer).GetThreadRuns(m, &assistantServiceGetThreadRunsServer{stream})
}

type AssistantService_GetThreadRunsServer interface {
	Send(*ThreadRun) error
	grpc.ServerStream
}

type assistantServiceGetThreadRunsServer struct {
	grpc.ServerStream
}

func (x *assistantServiceGetThreadRunsServer) Send(m *ThreadRun) error {
	return x.ServerStream.SendMsg(m)
}

// AssistantService_ServiceDesc is the grpc.ServiceDesc for AssistantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssistantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assistants.AssistantService",
	HandlerType: (*AssistantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAssistantFiles",
			Handler:    _AssistantService_ListAssistantFiles_Handler,
		},
		{
			MethodName: "DeleteAssistantFile",
			Handler:    _AssistantService_DeleteAssistantFile_Handler,
		},
		{
			MethodName: "LinkFileToAssistant",
			Handler:    _AssistantService_LinkFileToAssistant_Handler,
		},
		{
			MethodName: "CreateThread",
			Handler:    _AssistantService_CreateThread_Handler,
		},
		{
			MethodName: "GetThread",
			Handler:    _AssistantService_GetThread_Handler,
		},
		{
			MethodName: "DeleteThread",
			Handler:    _AssistantService_DeleteThread_Handler,
		},
		{
			MethodName: "CreateThreadRun",
			Handler:    _AssistantService_CreateThreadRun_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetThreadRuns",
			Handler:       _AssistantService_GetThreadRuns_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/assistants/assistants.proto",
}
