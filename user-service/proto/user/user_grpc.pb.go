// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/user/user.proto

package user

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// GetUser takes in an empty message as username should be taken via jwt metadata
	GetUser(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*UserResponse, error)
	GetMatchUsers(ctx context.Context, in *MatchUsersRequest, opts ...grpc.CallOption) (*MultipleUserResponse, error)
	EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetCosmetics(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*CosmeticsResponse, error)
	GetCosmetic(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*CosmeticResponse, error)
	GetCosmeticType(ctx context.Context, in *CosmeticTypeRequest, opts ...grpc.CallOption) (*CosmeticsResponse, error)
	PurchaseCosmetic(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*IDRequest, error)
	GetInventory(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*InventoryItemsResponse, error)
	GetInventoryItem(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*InventoryItemResponse, error)
	EquipInventoryItem(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*InventoryItemResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMatchUsers(ctx context.Context, in *MatchUsersRequest, opts ...grpc.CallOption) (*MultipleUserResponse, error) {
	out := new(MultipleUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetMatchUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/EditUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCosmetics(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*CosmeticsResponse, error) {
	out := new(CosmeticsResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetCosmetics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCosmetic(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*CosmeticResponse, error) {
	out := new(CosmeticResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetCosmetic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCosmeticType(ctx context.Context, in *CosmeticTypeRequest, opts ...grpc.CallOption) (*CosmeticsResponse, error) {
	out := new(CosmeticsResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetCosmeticType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PurchaseCosmetic(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*IDRequest, error) {
	out := new(IDRequest)
	err := c.cc.Invoke(ctx, "/user.UserService/PurchaseCosmetic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetInventory(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*InventoryItemsResponse, error) {
	out := new(InventoryItemsResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetInventoryItem(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*InventoryItemResponse, error) {
	out := new(InventoryItemResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetInventoryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EquipInventoryItem(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*InventoryItemResponse, error) {
	out := new(InventoryItemResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/EquipInventoryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// GetUser takes in an empty message as username should be taken via jwt metadata
	GetUser(context.Context, *EmptyMessage) (*UserResponse, error)
	GetMatchUsers(context.Context, *MatchUsersRequest) (*MultipleUserResponse, error)
	EditUser(context.Context, *EditUserRequest) (*UserResponse, error)
	GetCosmetics(context.Context, *EmptyMessage) (*CosmeticsResponse, error)
	GetCosmetic(context.Context, *IDRequest) (*CosmeticResponse, error)
	GetCosmeticType(context.Context, *CosmeticTypeRequest) (*CosmeticsResponse, error)
	PurchaseCosmetic(context.Context, *IDRequest) (*IDRequest, error)
	GetInventory(context.Context, *EmptyMessage) (*InventoryItemsResponse, error)
	GetInventoryItem(context.Context, *IDRequest) (*InventoryItemResponse, error)
	EquipInventoryItem(context.Context, *IDRequest) (*InventoryItemResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *EmptyMessage) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetMatchUsers(context.Context, *MatchUsersRequest) (*MultipleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchUsers not implemented")
}
func (UnimplementedUserServiceServer) EditUser(context.Context, *EditUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}
func (UnimplementedUserServiceServer) GetCosmetics(context.Context, *EmptyMessage) (*CosmeticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCosmetics not implemented")
}
func (UnimplementedUserServiceServer) GetCosmetic(context.Context, *IDRequest) (*CosmeticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCosmetic not implemented")
}
func (UnimplementedUserServiceServer) GetCosmeticType(context.Context, *CosmeticTypeRequest) (*CosmeticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCosmeticType not implemented")
}
func (UnimplementedUserServiceServer) PurchaseCosmetic(context.Context, *IDRequest) (*IDRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseCosmetic not implemented")
}
func (UnimplementedUserServiceServer) GetInventory(context.Context, *EmptyMessage) (*InventoryItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventory not implemented")
}
func (UnimplementedUserServiceServer) GetInventoryItem(context.Context, *IDRequest) (*InventoryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventoryItem not implemented")
}
func (UnimplementedUserServiceServer) EquipInventoryItem(context.Context, *IDRequest) (*InventoryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EquipInventoryItem not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMatchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMatchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetMatchUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMatchUsers(ctx, req.(*MatchUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/EditUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditUser(ctx, req.(*EditUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCosmetics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCosmetics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetCosmetics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCosmetics(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCosmetic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCosmetic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetCosmetic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCosmetic(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCosmeticType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CosmeticTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCosmeticType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetCosmeticType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCosmeticType(ctx, req.(*CosmeticTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PurchaseCosmetic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PurchaseCosmetic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/PurchaseCosmetic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PurchaseCosmetic(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetInventory(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetInventoryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetInventoryItem(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EquipInventoryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EquipInventoryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/EquipInventoryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EquipInventoryItem(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetMatchUsers",
			Handler:    _UserService_GetMatchUsers_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _UserService_EditUser_Handler,
		},
		{
			MethodName: "GetCosmetics",
			Handler:    _UserService_GetCosmetics_Handler,
		},
		{
			MethodName: "GetCosmetic",
			Handler:    _UserService_GetCosmetic_Handler,
		},
		{
			MethodName: "GetCosmeticType",
			Handler:    _UserService_GetCosmeticType_Handler,
		},
		{
			MethodName: "PurchaseCosmetic",
			Handler:    _UserService_PurchaseCosmetic_Handler,
		},
		{
			MethodName: "GetInventory",
			Handler:    _UserService_GetInventory_Handler,
		},
		{
			MethodName: "GetInventoryItem",
			Handler:    _UserService_GetInventoryItem_Handler,
		},
		{
			MethodName: "EquipInventoryItem",
			Handler:    _UserService_EquipInventoryItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/user.proto",
}
