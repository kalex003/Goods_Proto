// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: goods/goods.proto

package goods1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Goods_Insert_FullMethodName      = "/goods.Goods/Insert"
	Goods_Update_FullMethodName      = "/goods.Goods/Update"
	Goods_GetById_FullMethodName     = "/goods.Goods/GetById"
	Goods_GetByPlace_FullMethodName  = "/goods.Goods/GetByPlace"
	Goods_GetByTare_FullMethodName   = "/goods.Goods/GetByTare"
	Goods_GetHistory_FullMethodName  = "/goods.Goods/GetHistory"
	Goods_UpdateIsDel_FullMethodName = "/goods.Goods/UpdateIsDel"
)

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Auth is service for managing permissions and roles.
type GoodsClient interface {
	// Register registers a new user.
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	// Login logs in a user and returns an auth token.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// IsAdmin checks whether a user is an admin.
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*OneGetResponse, error)
	GetByPlace(ctx context.Context, in *GetByPlaceRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetByTare(ctx context.Context, in *GetByTareRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetHistory(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetResponse, error)
	UpdateIsDel(ctx context.Context, in *UpdateIsDelRequest, opts ...grpc.CallOption) (*UpdateIsDelResponse, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, Goods_Insert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, Goods_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*OneGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OneGetResponse)
	err := c.cc.Invoke(ctx, Goods_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetByPlace(ctx context.Context, in *GetByPlaceRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Goods_GetByPlace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetByTare(ctx context.Context, in *GetByTareRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Goods_GetByTare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetHistory(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Goods_GetHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateIsDel(ctx context.Context, in *UpdateIsDelRequest, opts ...grpc.CallOption) (*UpdateIsDelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateIsDelResponse)
	err := c.cc.Invoke(ctx, Goods_UpdateIsDel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility.
//
// Auth is service for managing permissions and roles.
type GoodsServer interface {
	// Register registers a new user.
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)
	// Login logs in a user and returns an auth token.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// IsAdmin checks whether a user is an admin.
	GetById(context.Context, *GetByIdRequest) (*OneGetResponse, error)
	GetByPlace(context.Context, *GetByPlaceRequest) (*GetResponse, error)
	GetByTare(context.Context, *GetByTareRequest) (*GetResponse, error)
	GetHistory(context.Context, *GetByIdRequest) (*GetResponse, error)
	UpdateIsDel(context.Context, *UpdateIsDelRequest) (*UpdateIsDelResponse, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGoodsServer struct{}

func (UnimplementedGoodsServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedGoodsServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGoodsServer) GetById(context.Context, *GetByIdRequest) (*OneGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedGoodsServer) GetByPlace(context.Context, *GetByPlaceRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPlace not implemented")
}
func (UnimplementedGoodsServer) GetByTare(context.Context, *GetByTareRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByTare not implemented")
}
func (UnimplementedGoodsServer) GetHistory(context.Context, *GetByIdRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedGoodsServer) UpdateIsDel(context.Context, *UpdateIsDelRequest) (*UpdateIsDelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIsDel not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}
func (UnimplementedGoodsServer) testEmbeddedByValue()               {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	// If the following call pancis, it indicates UnimplementedGoodsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetByPlace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByPlaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetByPlace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetByPlace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetByPlace(ctx, req.(*GetByPlaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetByTare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetByTare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetByTare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetByTare(ctx, req.(*GetByTareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetHistory(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateIsDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIsDelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateIsDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_UpdateIsDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateIsDel(ctx, req.(*UpdateIsDelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Goods_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Goods_Update_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _Goods_GetById_Handler,
		},
		{
			MethodName: "GetByPlace",
			Handler:    _Goods_GetByPlace_Handler,
		},
		{
			MethodName: "GetByTare",
			Handler:    _Goods_GetByTare_Handler,
		},
		{
			MethodName: "GetHistory",
			Handler:    _Goods_GetHistory_Handler,
		},
		{
			MethodName: "UpdateIsDel",
			Handler:    _Goods_UpdateIsDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods/goods.proto",
}
