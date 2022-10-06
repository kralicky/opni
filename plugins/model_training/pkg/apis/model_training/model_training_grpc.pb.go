// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/model_training/pkg/model_training/model_training.proto

package model_training

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ModelTrainingClient is the client API for ModelTraining service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelTrainingClient interface {
	WorkloadLogCount(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*WorkloadsList, error)
	ModelStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Reference, error)
	ModelTrainingParameters(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WorkloadsList, error)
}

type modelTrainingClient struct {
	cc grpc.ClientConnInterface
}

func NewModelTrainingClient(cc grpc.ClientConnInterface) ModelTrainingClient {
	return &modelTrainingClient{cc}
}

func (c *modelTrainingClient) WorkloadLogCount(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*WorkloadsList, error) {
	out := new(WorkloadsList)
	err := c.cc.Invoke(ctx, "/model_training.ModelTraining/WorkloadLogCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelTrainingClient) ModelStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Reference, error) {
	out := new(v1.Reference)
	err := c.cc.Invoke(ctx, "/model_training.ModelTraining/ModelStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelTrainingClient) ModelTrainingParameters(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WorkloadsList, error) {
	out := new(WorkloadsList)
	err := c.cc.Invoke(ctx, "/model_training.ModelTraining/ModelTrainingParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelTrainingServer is the server API for ModelTraining service.
// All implementations must embed UnimplementedModelTrainingServer
// for forward compatibility
type ModelTrainingServer interface {
	WorkloadLogCount(context.Context, *v1.Reference) (*WorkloadsList, error)
	ModelStatus(context.Context, *emptypb.Empty) (*v1.Reference, error)
	ModelTrainingParameters(context.Context, *emptypb.Empty) (*WorkloadsList, error)
	mustEmbedUnimplementedModelTrainingServer()
}

// UnimplementedModelTrainingServer must be embedded to have forward compatible implementations.
type UnimplementedModelTrainingServer struct {
}

func (UnimplementedModelTrainingServer) WorkloadLogCount(context.Context, *v1.Reference) (*WorkloadsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkloadLogCount not implemented")
}
func (UnimplementedModelTrainingServer) ModelStatus(context.Context, *emptypb.Empty) (*v1.Reference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelStatus not implemented")
}
func (UnimplementedModelTrainingServer) ModelTrainingParameters(context.Context, *emptypb.Empty) (*WorkloadsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelTrainingParameters not implemented")
}
func (UnimplementedModelTrainingServer) mustEmbedUnimplementedModelTrainingServer() {}

// UnsafeModelTrainingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelTrainingServer will
// result in compilation errors.
type UnsafeModelTrainingServer interface {
	mustEmbedUnimplementedModelTrainingServer()
}

func RegisterModelTrainingServer(s grpc.ServiceRegistrar, srv ModelTrainingServer) {
	s.RegisterService(&ModelTraining_ServiceDesc, srv)
}

func _ModelTraining_WorkloadLogCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).WorkloadLogCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_training.ModelTraining/WorkloadLogCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).WorkloadLogCount(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelTraining_ModelStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).ModelStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_training.ModelTraining/ModelStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).ModelStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelTraining_ModelTrainingParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).ModelTrainingParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_training.ModelTraining/ModelTrainingParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).ModelTrainingParameters(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelTraining_ServiceDesc is the grpc.ServiceDesc for ModelTraining service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelTraining_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model_training.ModelTraining",
	HandlerType: (*ModelTrainingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WorkloadLogCount",
			Handler:    _ModelTraining_WorkloadLogCount_Handler,
		},
		{
			MethodName: "ModelStatus",
			Handler:    _ModelTraining_ModelStatus_Handler,
		},
		{
			MethodName: "ModelTrainingParameters",
			Handler:    _ModelTraining_ModelTrainingParameters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/model_training/pkg/model_training/model_training.proto",
}