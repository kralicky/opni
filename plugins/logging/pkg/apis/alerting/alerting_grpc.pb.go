// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/logging/pkg/apis/alerting/alerting.proto

package alerting

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

const (
	MonitorManagement_CreateMonitor_FullMethodName = "/alerting.logging.MonitorManagement/CreateMonitor"
	MonitorManagement_GetMonitor_FullMethodName    = "/alerting.logging.MonitorManagement/GetMonitor"
	MonitorManagement_UpdateMonitor_FullMethodName = "/alerting.logging.MonitorManagement/UpdateMonitor"
	MonitorManagement_DeleteMonitor_FullMethodName = "/alerting.logging.MonitorManagement/DeleteMonitor"
)

// MonitorManagementClient is the client API for MonitorManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitorManagementClient interface {
	CreateMonitor(ctx context.Context, in *Monitor, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMonitor(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*Monitor, error)
	UpdateMonitor(ctx context.Context, in *Monitor, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteMonitor(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type monitorManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorManagementClient(cc grpc.ClientConnInterface) MonitorManagementClient {
	return &monitorManagementClient{cc}
}

func (c *monitorManagementClient) CreateMonitor(ctx context.Context, in *Monitor, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MonitorManagement_CreateMonitor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorManagementClient) GetMonitor(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*Monitor, error) {
	out := new(Monitor)
	err := c.cc.Invoke(ctx, MonitorManagement_GetMonitor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorManagementClient) UpdateMonitor(ctx context.Context, in *Monitor, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MonitorManagement_UpdateMonitor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorManagementClient) DeleteMonitor(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MonitorManagement_DeleteMonitor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorManagementServer is the server API for MonitorManagement service.
// All implementations must embed UnimplementedMonitorManagementServer
// for forward compatibility
type MonitorManagementServer interface {
	CreateMonitor(context.Context, *Monitor) (*emptypb.Empty, error)
	GetMonitor(context.Context, *v1.Reference) (*Monitor, error)
	UpdateMonitor(context.Context, *Monitor) (*emptypb.Empty, error)
	DeleteMonitor(context.Context, *v1.Reference) (*emptypb.Empty, error)
	mustEmbedUnimplementedMonitorManagementServer()
}

// UnimplementedMonitorManagementServer must be embedded to have forward compatible implementations.
type UnimplementedMonitorManagementServer struct {
}

func (UnimplementedMonitorManagementServer) CreateMonitor(context.Context, *Monitor) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMonitor not implemented")
}
func (UnimplementedMonitorManagementServer) GetMonitor(context.Context, *v1.Reference) (*Monitor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonitor not implemented")
}
func (UnimplementedMonitorManagementServer) UpdateMonitor(context.Context, *Monitor) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMonitor not implemented")
}
func (UnimplementedMonitorManagementServer) DeleteMonitor(context.Context, *v1.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMonitor not implemented")
}
func (UnimplementedMonitorManagementServer) mustEmbedUnimplementedMonitorManagementServer() {}

// UnsafeMonitorManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitorManagementServer will
// result in compilation errors.
type UnsafeMonitorManagementServer interface {
	mustEmbedUnimplementedMonitorManagementServer()
}

func RegisterMonitorManagementServer(s grpc.ServiceRegistrar, srv MonitorManagementServer) {
	s.RegisterService(&MonitorManagement_ServiceDesc, srv)
}

func _MonitorManagement_CreateMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Monitor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorManagementServer).CreateMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorManagement_CreateMonitor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorManagementServer).CreateMonitor(ctx, req.(*Monitor))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorManagement_GetMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorManagementServer).GetMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorManagement_GetMonitor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorManagementServer).GetMonitor(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorManagement_UpdateMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Monitor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorManagementServer).UpdateMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorManagement_UpdateMonitor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorManagementServer).UpdateMonitor(ctx, req.(*Monitor))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorManagement_DeleteMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorManagementServer).DeleteMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorManagement_DeleteMonitor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorManagementServer).DeleteMonitor(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

// MonitorManagement_ServiceDesc is the grpc.ServiceDesc for MonitorManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonitorManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alerting.logging.MonitorManagement",
	HandlerType: (*MonitorManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMonitor",
			Handler:    _MonitorManagement_CreateMonitor_Handler,
		},
		{
			MethodName: "GetMonitor",
			Handler:    _MonitorManagement_GetMonitor_Handler,
		},
		{
			MethodName: "UpdateMonitor",
			Handler:    _MonitorManagement_UpdateMonitor_Handler,
		},
		{
			MethodName: "DeleteMonitor",
			Handler:    _MonitorManagement_DeleteMonitor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/logging/pkg/apis/alerting/alerting.proto",
}

const (
	NotificationManagement_CreateNotification_FullMethodName = "/alerting.logging.NotificationManagement/CreateNotification"
	NotificationManagement_GetNotification_FullMethodName    = "/alerting.logging.NotificationManagement/GetNotification"
	NotificationManagement_ListNotifications_FullMethodName  = "/alerting.logging.NotificationManagement/ListNotifications"
	NotificationManagement_UpdateNotification_FullMethodName = "/alerting.logging.NotificationManagement/UpdateNotification"
	NotificationManagement_DeleteDestination_FullMethodName  = "/alerting.logging.NotificationManagement/DeleteDestination"
)

// NotificationManagementClient is the client API for NotificationManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationManagementClient interface {
	CreateNotification(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetNotification(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListNotifications(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChannelList, error)
	UpdateNotification(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteDestination(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type notificationManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationManagementClient(cc grpc.ClientConnInterface) NotificationManagementClient {
	return &notificationManagementClient{cc}
}

func (c *notificationManagementClient) CreateNotification(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NotificationManagement_CreateNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagementClient) GetNotification(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NotificationManagement_GetNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagementClient) ListNotifications(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChannelList, error) {
	out := new(ChannelList)
	err := c.cc.Invoke(ctx, NotificationManagement_ListNotifications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagementClient) UpdateNotification(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NotificationManagement_UpdateNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationManagementClient) DeleteDestination(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NotificationManagement_DeleteDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationManagementServer is the server API for NotificationManagement service.
// All implementations must embed UnimplementedNotificationManagementServer
// for forward compatibility
type NotificationManagementServer interface {
	CreateNotification(context.Context, *Channel) (*emptypb.Empty, error)
	GetNotification(context.Context, *v1.Reference) (*emptypb.Empty, error)
	ListNotifications(context.Context, *emptypb.Empty) (*ChannelList, error)
	UpdateNotification(context.Context, *Channel) (*emptypb.Empty, error)
	DeleteDestination(context.Context, *v1.Reference) (*emptypb.Empty, error)
	mustEmbedUnimplementedNotificationManagementServer()
}

// UnimplementedNotificationManagementServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationManagementServer struct {
}

func (UnimplementedNotificationManagementServer) CreateNotification(context.Context, *Channel) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedNotificationManagementServer) GetNotification(context.Context, *v1.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedNotificationManagementServer) ListNotifications(context.Context, *emptypb.Empty) (*ChannelList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotifications not implemented")
}
func (UnimplementedNotificationManagementServer) UpdateNotification(context.Context, *Channel) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotification not implemented")
}
func (UnimplementedNotificationManagementServer) DeleteDestination(context.Context, *v1.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDestination not implemented")
}
func (UnimplementedNotificationManagementServer) mustEmbedUnimplementedNotificationManagementServer() {
}

// UnsafeNotificationManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationManagementServer will
// result in compilation errors.
type UnsafeNotificationManagementServer interface {
	mustEmbedUnimplementedNotificationManagementServer()
}

func RegisterNotificationManagementServer(s grpc.ServiceRegistrar, srv NotificationManagementServer) {
	s.RegisterService(&NotificationManagement_ServiceDesc, srv)
}

func _NotificationManagement_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagementServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationManagement_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagementServer).CreateNotification(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManagement_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagementServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationManagement_GetNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagementServer).GetNotification(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManagement_ListNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagementServer).ListNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationManagement_ListNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagementServer).ListNotifications(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManagement_UpdateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagementServer).UpdateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationManagement_UpdateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagementServer).UpdateNotification(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationManagement_DeleteDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationManagementServer).DeleteDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationManagement_DeleteDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationManagementServer).DeleteDestination(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationManagement_ServiceDesc is the grpc.ServiceDesc for NotificationManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alerting.logging.NotificationManagement",
	HandlerType: (*NotificationManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotification",
			Handler:    _NotificationManagement_CreateNotification_Handler,
		},
		{
			MethodName: "GetNotification",
			Handler:    _NotificationManagement_GetNotification_Handler,
		},
		{
			MethodName: "ListNotifications",
			Handler:    _NotificationManagement_ListNotifications_Handler,
		},
		{
			MethodName: "UpdateNotification",
			Handler:    _NotificationManagement_UpdateNotification_Handler,
		},
		{
			MethodName: "DeleteDestination",
			Handler:    _NotificationManagement_DeleteDestination_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/logging/pkg/apis/alerting/alerting.proto",
}

const (
	AlertManagement_ListAlerts_FullMethodName       = "/alerting.logging.AlertManagement/ListAlerts"
	AlertManagement_AcknowledgeAlert_FullMethodName = "/alerting.logging.AlertManagement/AcknowledgeAlert"
)

// AlertManagementClient is the client API for AlertManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertManagementClient interface {
	ListAlerts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAlertsResponse, error)
	AcknowledgeAlert(ctx context.Context, in *AcknowledgeAlertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type alertManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertManagementClient(cc grpc.ClientConnInterface) AlertManagementClient {
	return &alertManagementClient{cc}
}

func (c *alertManagementClient) ListAlerts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAlertsResponse, error) {
	out := new(ListAlertsResponse)
	err := c.cc.Invoke(ctx, AlertManagement_ListAlerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertManagementClient) AcknowledgeAlert(ctx context.Context, in *AcknowledgeAlertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AlertManagement_AcknowledgeAlert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertManagementServer is the server API for AlertManagement service.
// All implementations must embed UnimplementedAlertManagementServer
// for forward compatibility
type AlertManagementServer interface {
	ListAlerts(context.Context, *emptypb.Empty) (*ListAlertsResponse, error)
	AcknowledgeAlert(context.Context, *AcknowledgeAlertRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAlertManagementServer()
}

// UnimplementedAlertManagementServer must be embedded to have forward compatible implementations.
type UnimplementedAlertManagementServer struct {
}

func (UnimplementedAlertManagementServer) ListAlerts(context.Context, *emptypb.Empty) (*ListAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlerts not implemented")
}
func (UnimplementedAlertManagementServer) AcknowledgeAlert(context.Context, *AcknowledgeAlertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcknowledgeAlert not implemented")
}
func (UnimplementedAlertManagementServer) mustEmbedUnimplementedAlertManagementServer() {}

// UnsafeAlertManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertManagementServer will
// result in compilation errors.
type UnsafeAlertManagementServer interface {
	mustEmbedUnimplementedAlertManagementServer()
}

func RegisterAlertManagementServer(s grpc.ServiceRegistrar, srv AlertManagementServer) {
	s.RegisterService(&AlertManagement_ServiceDesc, srv)
}

func _AlertManagement_ListAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertManagementServer).ListAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertManagement_ListAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertManagementServer).ListAlerts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertManagement_AcknowledgeAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertManagementServer).AcknowledgeAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertManagement_AcknowledgeAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertManagementServer).AcknowledgeAlert(ctx, req.(*AcknowledgeAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertManagement_ServiceDesc is the grpc.ServiceDesc for AlertManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alerting.logging.AlertManagement",
	HandlerType: (*AlertManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAlerts",
			Handler:    _AlertManagement_ListAlerts_Handler,
		},
		{
			MethodName: "AcknowledgeAlert",
			Handler:    _AlertManagement_AcknowledgeAlert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/logging/pkg/apis/alerting/alerting.proto",
}