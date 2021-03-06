// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package configmodel

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

// ModelManageClient is the client API for ModelManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelManageClient interface {
	SaveModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*OpRet, error)
	DeleteModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*OpRet, error)
	UpdateModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*OpRet, error)
	FindModels(ctx context.Context, in *Filter, opts ...grpc.CallOption) (ModelManage_FindModelsClient, error)
}

type modelManageClient struct {
	cc grpc.ClientConnInterface
}

func NewModelManageClient(cc grpc.ClientConnInterface) ModelManageClient {
	return &modelManageClient{cc}
}

func (c *modelManageClient) SaveModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*OpRet, error) {
	out := new(OpRet)
	err := c.cc.Invoke(ctx, "/configmodel.ModelManage/SaveModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelManageClient) DeleteModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*OpRet, error) {
	out := new(OpRet)
	err := c.cc.Invoke(ctx, "/configmodel.ModelManage/DeleteModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelManageClient) UpdateModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*OpRet, error) {
	out := new(OpRet)
	err := c.cc.Invoke(ctx, "/configmodel.ModelManage/UpdateModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelManageClient) FindModels(ctx context.Context, in *Filter, opts ...grpc.CallOption) (ModelManage_FindModelsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ModelManage_ServiceDesc.Streams[0], "/configmodel.ModelManage/FindModels", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelManageFindModelsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ModelManage_FindModelsClient interface {
	Recv() (*Model, error)
	grpc.ClientStream
}

type modelManageFindModelsClient struct {
	grpc.ClientStream
}

func (x *modelManageFindModelsClient) Recv() (*Model, error) {
	m := new(Model)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ModelManageServer is the server API for ModelManage service.
// All implementations must embed UnimplementedModelManageServer
// for forward compatibility
type ModelManageServer interface {
	SaveModel(context.Context, *Model) (*OpRet, error)
	DeleteModel(context.Context, *Model) (*OpRet, error)
	UpdateModel(context.Context, *Model) (*OpRet, error)
	FindModels(*Filter, ModelManage_FindModelsServer) error
	mustEmbedUnimplementedModelManageServer()
}

// UnimplementedModelManageServer must be embedded to have forward compatible implementations.
type UnimplementedModelManageServer struct {
}

func (UnimplementedModelManageServer) SaveModel(context.Context, *Model) (*OpRet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveModel not implemented")
}
func (UnimplementedModelManageServer) DeleteModel(context.Context, *Model) (*OpRet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModel not implemented")
}
func (UnimplementedModelManageServer) UpdateModel(context.Context, *Model) (*OpRet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModel not implemented")
}
func (UnimplementedModelManageServer) FindModels(*Filter, ModelManage_FindModelsServer) error {
	return status.Errorf(codes.Unimplemented, "method FindModels not implemented")
}
func (UnimplementedModelManageServer) mustEmbedUnimplementedModelManageServer() {}

// UnsafeModelManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelManageServer will
// result in compilation errors.
type UnsafeModelManageServer interface {
	mustEmbedUnimplementedModelManageServer()
}

func RegisterModelManageServer(s grpc.ServiceRegistrar, srv ModelManageServer) {
	s.RegisterService(&ModelManage_ServiceDesc, srv)
}

func _ModelManage_SaveModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Model)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelManageServer).SaveModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configmodel.ModelManage/SaveModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelManageServer).SaveModel(ctx, req.(*Model))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelManage_DeleteModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Model)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelManageServer).DeleteModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configmodel.ModelManage/DeleteModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelManageServer).DeleteModel(ctx, req.(*Model))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelManage_UpdateModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Model)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelManageServer).UpdateModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configmodel.ModelManage/UpdateModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelManageServer).UpdateModel(ctx, req.(*Model))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelManage_FindModels_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Filter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModelManageServer).FindModels(m, &modelManageFindModelsServer{stream})
}

type ModelManage_FindModelsServer interface {
	Send(*Model) error
	grpc.ServerStream
}

type modelManageFindModelsServer struct {
	grpc.ServerStream
}

func (x *modelManageFindModelsServer) Send(m *Model) error {
	return x.ServerStream.SendMsg(m)
}

// ModelManage_ServiceDesc is the grpc.ServiceDesc for ModelManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "configmodel.ModelManage",
	HandlerType: (*ModelManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveModel",
			Handler:    _ModelManage_SaveModel_Handler,
		},
		{
			MethodName: "DeleteModel",
			Handler:    _ModelManage_DeleteModel_Handler,
		},
		{
			MethodName: "UpdateModel",
			Handler:    _ModelManage_UpdateModel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindModels",
			Handler:       _ModelManage_FindModels_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "configmodel-model.proto",
}
