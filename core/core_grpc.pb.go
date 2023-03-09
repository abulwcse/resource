// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: core/core.proto

package core

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

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreClient interface {
	GetCustomers(ctx context.Context, in *EmptyParamRequest, opts ...grpc.CallOption) (*StandardReply, error)
	GetCustomerProject(ctx context.Context, in *GetCustomerProjectRequest, opts ...grpc.CallOption) (*StandardReply, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*StandardReply, error)
	GetChart(ctx context.Context, in *GetChartRequest, opts ...grpc.CallOption) (*StandardReply, error)
	GetFilters(ctx context.Context, in *GetFiltersRequest, opts ...grpc.CallOption) (*StandardReply, error)
	GetMarkers(ctx context.Context, in *GetMarkersRequest, opts ...grpc.CallOption) (*StandardReply, error)
	GetTaxonomySummary(ctx context.Context, in *GetTaxonomySummaryRequest, opts ...grpc.CallOption) (*StandardReply, error)
	GetSamples(ctx context.Context, in *GetSamplesRequest, opts ...grpc.CallOption) (*StandardReply, error)
}

type coreClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreClient(cc grpc.ClientConnInterface) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) GetCustomers(ctx context.Context, in *EmptyParamRequest, opts ...grpc.CallOption) (*StandardReply, error) {
	out := new(StandardReply)
	err := c.cc.Invoke(ctx, "/core.Core/GetCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetCustomerProject(ctx context.Context, in *GetCustomerProjectRequest, opts ...grpc.CallOption) (*StandardReply, error) {
	out := new(StandardReply)
	err := c.cc.Invoke(ctx, "/core.Core/GetCustomerProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*StandardReply, error) {
	out := new(StandardReply)
	err := c.cc.Invoke(ctx, "/core.Core/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetChart(ctx context.Context, in *GetChartRequest, opts ...grpc.CallOption) (*StandardReply, error) {
	out := new(StandardReply)
	err := c.cc.Invoke(ctx, "/core.Core/GetChart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetFilters(ctx context.Context, in *GetFiltersRequest, opts ...grpc.CallOption) (*StandardReply, error) {
	out := new(StandardReply)
	err := c.cc.Invoke(ctx, "/core.Core/GetFilters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetMarkers(ctx context.Context, in *GetMarkersRequest, opts ...grpc.CallOption) (*StandardReply, error) {
	out := new(StandardReply)
	err := c.cc.Invoke(ctx, "/core.Core/GetMarkers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetTaxonomySummary(ctx context.Context, in *GetTaxonomySummaryRequest, opts ...grpc.CallOption) (*StandardReply, error) {
	out := new(StandardReply)
	err := c.cc.Invoke(ctx, "/core.Core/GetTaxonomySummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetSamples(ctx context.Context, in *GetSamplesRequest, opts ...grpc.CallOption) (*StandardReply, error) {
	out := new(StandardReply)
	err := c.cc.Invoke(ctx, "/core.Core/GetSamples", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
// All implementations must embed UnimplementedCoreServer
// for forward compatibility
type CoreServer interface {
	GetCustomers(context.Context, *EmptyParamRequest) (*StandardReply, error)
	GetCustomerProject(context.Context, *GetCustomerProjectRequest) (*StandardReply, error)
	GetProject(context.Context, *GetProjectRequest) (*StandardReply, error)
	GetChart(context.Context, *GetChartRequest) (*StandardReply, error)
	GetFilters(context.Context, *GetFiltersRequest) (*StandardReply, error)
	GetMarkers(context.Context, *GetMarkersRequest) (*StandardReply, error)
	GetTaxonomySummary(context.Context, *GetTaxonomySummaryRequest) (*StandardReply, error)
	GetSamples(context.Context, *GetSamplesRequest) (*StandardReply, error)
	mustEmbedUnimplementedCoreServer()
}

// UnimplementedCoreServer must be embedded to have forward compatible implementations.
type UnimplementedCoreServer struct {
}

func (UnimplementedCoreServer) GetCustomers(context.Context, *EmptyParamRequest) (*StandardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomers not implemented")
}
func (UnimplementedCoreServer) GetCustomerProject(context.Context, *GetCustomerProjectRequest) (*StandardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerProject not implemented")
}
func (UnimplementedCoreServer) GetProject(context.Context, *GetProjectRequest) (*StandardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedCoreServer) GetChart(context.Context, *GetChartRequest) (*StandardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChart not implemented")
}
func (UnimplementedCoreServer) GetFilters(context.Context, *GetFiltersRequest) (*StandardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilters not implemented")
}
func (UnimplementedCoreServer) GetMarkers(context.Context, *GetMarkersRequest) (*StandardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarkers not implemented")
}
func (UnimplementedCoreServer) GetTaxonomySummary(context.Context, *GetTaxonomySummaryRequest) (*StandardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaxonomySummary not implemented")
}
func (UnimplementedCoreServer) GetSamples(context.Context, *GetSamplesRequest) (*StandardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSamples not implemented")
}
func (UnimplementedCoreServer) mustEmbedUnimplementedCoreServer() {}

// UnsafeCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServer will
// result in compilation errors.
type UnsafeCoreServer interface {
	mustEmbedUnimplementedCoreServer()
}

func RegisterCoreServer(s grpc.ServiceRegistrar, srv CoreServer) {
	s.RegisterService(&Core_ServiceDesc, srv)
}

func _Core_GetCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Core/GetCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetCustomers(ctx, req.(*EmptyParamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetCustomerProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetCustomerProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Core/GetCustomerProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetCustomerProject(ctx, req.(*GetCustomerProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Core/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Core/GetChart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetChart(ctx, req.(*GetChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Core/GetFilters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetFilters(ctx, req.(*GetFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetMarkers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarkersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetMarkers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Core/GetMarkers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetMarkers(ctx, req.(*GetMarkersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetTaxonomySummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaxonomySummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetTaxonomySummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Core/GetTaxonomySummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetTaxonomySummary(ctx, req.(*GetTaxonomySummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetSamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetSamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Core/GetSamples",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetSamples(ctx, req.(*GetSamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Core_ServiceDesc is the grpc.ServiceDesc for Core service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Core_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomers",
			Handler:    _Core_GetCustomers_Handler,
		},
		{
			MethodName: "GetCustomerProject",
			Handler:    _Core_GetCustomerProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Core_GetProject_Handler,
		},
		{
			MethodName: "GetChart",
			Handler:    _Core_GetChart_Handler,
		},
		{
			MethodName: "GetFilters",
			Handler:    _Core_GetFilters_Handler,
		},
		{
			MethodName: "GetMarkers",
			Handler:    _Core_GetMarkers_Handler,
		},
		{
			MethodName: "GetTaxonomySummary",
			Handler:    _Core_GetTaxonomySummary_Handler,
		},
		{
			MethodName: "GetSamples",
			Handler:    _Core_GetSamples_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/core.proto",
}
