// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: entity.proto

package entity

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

// EntityStoreClient is the client API for EntityStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntityStoreClient interface {
	Read(ctx context.Context, in *ReadEntityRequest, opts ...grpc.CallOption) (*Entity, error)
	BatchRead(ctx context.Context, in *BatchReadEntityRequest, opts ...grpc.CallOption) (*BatchReadEntityResponse, error)
	Write(ctx context.Context, in *WriteEntityRequest, opts ...grpc.CallOption) (*WriteEntityResponse, error)
	Delete(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteEntityResponse, error)
	History(ctx context.Context, in *EntityHistoryRequest, opts ...grpc.CallOption) (*EntityHistoryResponse, error)
	Search(ctx context.Context, in *EntitySearchRequest, opts ...grpc.CallOption) (*EntitySearchResponse, error)
	Watch(ctx context.Context, in *EntityWatchRequest, opts ...grpc.CallOption) (EntityStore_WatchClient, error)
	// TEMPORARY... while we split this into a new service (see below)
	AdminWrite(ctx context.Context, in *AdminWriteEntityRequest, opts ...grpc.CallOption) (*WriteEntityResponse, error)
}

type entityStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityStoreClient(cc grpc.ClientConnInterface) EntityStoreClient {
	return &entityStoreClient{cc}
}

func (c *entityStoreClient) Read(ctx context.Context, in *ReadEntityRequest, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) BatchRead(ctx context.Context, in *BatchReadEntityRequest, opts ...grpc.CallOption) (*BatchReadEntityResponse, error) {
	out := new(BatchReadEntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/BatchRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) Write(ctx context.Context, in *WriteEntityRequest, opts ...grpc.CallOption) (*WriteEntityResponse, error) {
	out := new(WriteEntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) Delete(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteEntityResponse, error) {
	out := new(DeleteEntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) History(ctx context.Context, in *EntityHistoryRequest, opts ...grpc.CallOption) (*EntityHistoryResponse, error) {
	out := new(EntityHistoryResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/History", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) Search(ctx context.Context, in *EntitySearchRequest, opts ...grpc.CallOption) (*EntitySearchResponse, error) {
	out := new(EntitySearchResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityStoreClient) Watch(ctx context.Context, in *EntityWatchRequest, opts ...grpc.CallOption) (EntityStore_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &EntityStore_ServiceDesc.Streams[0], "/entity.EntityStore/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &entityStoreWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EntityStore_WatchClient interface {
	Recv() (*EntityWatchResponse, error)
	grpc.ClientStream
}

type entityStoreWatchClient struct {
	grpc.ClientStream
}

func (x *entityStoreWatchClient) Recv() (*EntityWatchResponse, error) {
	m := new(EntityWatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *entityStoreClient) AdminWrite(ctx context.Context, in *AdminWriteEntityRequest, opts ...grpc.CallOption) (*WriteEntityResponse, error) {
	out := new(WriteEntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStore/AdminWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityStoreServer is the server API for EntityStore service.
// All implementations should embed UnimplementedEntityStoreServer
// for forward compatibility
type EntityStoreServer interface {
	Read(context.Context, *ReadEntityRequest) (*Entity, error)
	BatchRead(context.Context, *BatchReadEntityRequest) (*BatchReadEntityResponse, error)
	Write(context.Context, *WriteEntityRequest) (*WriteEntityResponse, error)
	Delete(context.Context, *DeleteEntityRequest) (*DeleteEntityResponse, error)
	History(context.Context, *EntityHistoryRequest) (*EntityHistoryResponse, error)
	Search(context.Context, *EntitySearchRequest) (*EntitySearchResponse, error)
	Watch(*EntityWatchRequest, EntityStore_WatchServer) error
	// TEMPORARY... while we split this into a new service (see below)
	AdminWrite(context.Context, *AdminWriteEntityRequest) (*WriteEntityResponse, error)
}

// UnimplementedEntityStoreServer should be embedded to have forward compatible implementations.
type UnimplementedEntityStoreServer struct {
}

func (UnimplementedEntityStoreServer) Read(context.Context, *ReadEntityRequest) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedEntityStoreServer) BatchRead(context.Context, *BatchReadEntityRequest) (*BatchReadEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchRead not implemented")
}
func (UnimplementedEntityStoreServer) Write(context.Context, *WriteEntityRequest) (*WriteEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedEntityStoreServer) Delete(context.Context, *DeleteEntityRequest) (*DeleteEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEntityStoreServer) History(context.Context, *EntityHistoryRequest) (*EntityHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (UnimplementedEntityStoreServer) Search(context.Context, *EntitySearchRequest) (*EntitySearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedEntityStoreServer) Watch(*EntityWatchRequest, EntityStore_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedEntityStoreServer) AdminWrite(context.Context, *AdminWriteEntityRequest) (*WriteEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminWrite not implemented")
}

// UnsafeEntityStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityStoreServer will
// result in compilation errors.
type UnsafeEntityStoreServer interface {
	mustEmbedUnimplementedEntityStoreServer()
}

func RegisterEntityStoreServer(s grpc.ServiceRegistrar, srv EntityStoreServer) {
	s.RegisterService(&EntityStore_ServiceDesc, srv)
}

func _EntityStore_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).Read(ctx, req.(*ReadEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_BatchRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchReadEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).BatchRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/BatchRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).BatchRead(ctx, req.(*BatchReadEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).Write(ctx, req.(*WriteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).Delete(ctx, req.(*DeleteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/History",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).History(ctx, req.(*EntityHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntitySearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).Search(ctx, req.(*EntitySearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityStore_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EntityWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EntityStoreServer).Watch(m, &entityStoreWatchServer{stream})
}

type EntityStore_WatchServer interface {
	Send(*EntityWatchResponse) error
	grpc.ServerStream
}

type entityStoreWatchServer struct {
	grpc.ServerStream
}

func (x *entityStoreWatchServer) Send(m *EntityWatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EntityStore_AdminWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminWriteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreServer).AdminWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStore/AdminWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreServer).AdminWrite(ctx, req.(*AdminWriteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntityStore_ServiceDesc is the grpc.ServiceDesc for EntityStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entity.EntityStore",
	HandlerType: (*EntityStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _EntityStore_Read_Handler,
		},
		{
			MethodName: "BatchRead",
			Handler:    _EntityStore_BatchRead_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _EntityStore_Write_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EntityStore_Delete_Handler,
		},
		{
			MethodName: "History",
			Handler:    _EntityStore_History_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _EntityStore_Search_Handler,
		},
		{
			MethodName: "AdminWrite",
			Handler:    _EntityStore_AdminWrite_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _EntityStore_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entity.proto",
}

// EntityStoreAdminClient is the client API for EntityStoreAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntityStoreAdminClient interface {
	AdminWrite(ctx context.Context, in *AdminWriteEntityRequest, opts ...grpc.CallOption) (*WriteEntityResponse, error)
}

type entityStoreAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityStoreAdminClient(cc grpc.ClientConnInterface) EntityStoreAdminClient {
	return &entityStoreAdminClient{cc}
}

func (c *entityStoreAdminClient) AdminWrite(ctx context.Context, in *AdminWriteEntityRequest, opts ...grpc.CallOption) (*WriteEntityResponse, error) {
	out := new(WriteEntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityStoreAdmin/AdminWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityStoreAdminServer is the server API for EntityStoreAdmin service.
// All implementations should embed UnimplementedEntityStoreAdminServer
// for forward compatibility
type EntityStoreAdminServer interface {
	AdminWrite(context.Context, *AdminWriteEntityRequest) (*WriteEntityResponse, error)
}

// UnimplementedEntityStoreAdminServer should be embedded to have forward compatible implementations.
type UnimplementedEntityStoreAdminServer struct {
}

func (UnimplementedEntityStoreAdminServer) AdminWrite(context.Context, *AdminWriteEntityRequest) (*WriteEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminWrite not implemented")
}

// UnsafeEntityStoreAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityStoreAdminServer will
// result in compilation errors.
type UnsafeEntityStoreAdminServer interface {
	mustEmbedUnimplementedEntityStoreAdminServer()
}

func RegisterEntityStoreAdminServer(s grpc.ServiceRegistrar, srv EntityStoreAdminServer) {
	s.RegisterService(&EntityStoreAdmin_ServiceDesc, srv)
}

func _EntityStoreAdmin_AdminWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminWriteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityStoreAdminServer).AdminWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityStoreAdmin/AdminWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityStoreAdminServer).AdminWrite(ctx, req.(*AdminWriteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntityStoreAdmin_ServiceDesc is the grpc.ServiceDesc for EntityStoreAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityStoreAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entity.EntityStoreAdmin",
	HandlerType: (*EntityStoreAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminWrite",
			Handler:    _EntityStoreAdmin_AdminWrite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entity.proto",
}