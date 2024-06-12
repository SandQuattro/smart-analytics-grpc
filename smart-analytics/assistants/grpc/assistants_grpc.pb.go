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
)

// AssistantServiceClient is the client API for AssistantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssistantServiceClient interface {
	ListAssistantFiles(ctx context.Context, in *AssistantRequest, opts ...grpc.CallOption) (*ListAssistantFilesResponse, error)
	DeleteAssistantFile(ctx context.Context, in *AssistantFileRequest, opts ...grpc.CallOption) (*DeletedObject, error)
	LinkFileToAssistant(ctx context.Context, in *AssistantFileRequest, opts ...grpc.CallOption) (*AssistantFileData, error)
	CreateThread(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AssistantObject, error)
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

// AssistantServiceServer is the server API for AssistantService service.
// All implementations must embed UnimplementedAssistantServiceServer
// for forward compatibility
type AssistantServiceServer interface {
	ListAssistantFiles(context.Context, *AssistantRequest) (*ListAssistantFilesResponse, error)
	DeleteAssistantFile(context.Context, *AssistantFileRequest) (*DeletedObject, error)
	LinkFileToAssistant(context.Context, *AssistantFileRequest) (*AssistantFileData, error)
	CreateThread(context.Context, *emptypb.Empty) (*AssistantObject, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/assistants/assistants.proto",
}
