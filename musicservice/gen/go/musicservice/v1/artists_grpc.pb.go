// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: musicservice/v1/artists.proto

package protos

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
	ArtistService_CreateArtist_FullMethodName = "/musicservice.ArtistService/CreateArtist"
	ArtistService_GetArtists_FullMethodName   = "/musicservice.ArtistService/GetArtists"
	ArtistService_UpdateArtist_FullMethodName = "/musicservice.ArtistService/UpdateArtist"
	ArtistService_DeleteArtist_FullMethodName = "/musicservice.ArtistService/DeleteArtist"
	ArtistService_ListArtists_FullMethodName  = "/musicservice.ArtistService/ListArtists"
)

// ArtistServiceClient is the client API for ArtistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtistServiceClient interface {
	// *
	// Create a new Artist
	CreateArtist(ctx context.Context, in *CreateArtistRequest, opts ...grpc.CallOption) (*CreateArtistResponse, error)
	// *
	// Batch gets multiple artists.
	GetArtists(ctx context.Context, in *GetArtistsRequest, opts ...grpc.CallOption) (*GetArtistsResponse, error)
	// *
	// Updates specific fields of an Artist
	UpdateArtist(ctx context.Context, in *UpdateArtistRequest, opts ...grpc.CallOption) (*UpdateArtistResponse, error)
	// *
	// Deletes an artist from our system.
	DeleteArtist(ctx context.Context, in *DeleteArtistRequest, opts ...grpc.CallOption) (*DeleteArtistResponse, error)
	// *
	// Finds and retrieves artists matching the particular criteria.
	ListArtists(ctx context.Context, in *ListArtistsRequest, opts ...grpc.CallOption) (*ListArtistsResponse, error)
}

type artistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArtistServiceClient(cc grpc.ClientConnInterface) ArtistServiceClient {
	return &artistServiceClient{cc}
}

func (c *artistServiceClient) CreateArtist(ctx context.Context, in *CreateArtistRequest, opts ...grpc.CallOption) (*CreateArtistResponse, error) {
	out := new(CreateArtistResponse)
	err := c.cc.Invoke(ctx, ArtistService_CreateArtist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) GetArtists(ctx context.Context, in *GetArtistsRequest, opts ...grpc.CallOption) (*GetArtistsResponse, error) {
	out := new(GetArtistsResponse)
	err := c.cc.Invoke(ctx, ArtistService_GetArtists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) UpdateArtist(ctx context.Context, in *UpdateArtistRequest, opts ...grpc.CallOption) (*UpdateArtistResponse, error) {
	out := new(UpdateArtistResponse)
	err := c.cc.Invoke(ctx, ArtistService_UpdateArtist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) DeleteArtist(ctx context.Context, in *DeleteArtistRequest, opts ...grpc.CallOption) (*DeleteArtistResponse, error) {
	out := new(DeleteArtistResponse)
	err := c.cc.Invoke(ctx, ArtistService_DeleteArtist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) ListArtists(ctx context.Context, in *ListArtistsRequest, opts ...grpc.CallOption) (*ListArtistsResponse, error) {
	out := new(ListArtistsResponse)
	err := c.cc.Invoke(ctx, ArtistService_ListArtists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtistServiceServer is the server API for ArtistService service.
// All implementations must embed UnimplementedArtistServiceServer
// for forward compatibility
type ArtistServiceServer interface {
	// *
	// Create a new Artist
	CreateArtist(context.Context, *CreateArtistRequest) (*CreateArtistResponse, error)
	// *
	// Batch gets multiple artists.
	GetArtists(context.Context, *GetArtistsRequest) (*GetArtistsResponse, error)
	// *
	// Updates specific fields of an Artist
	UpdateArtist(context.Context, *UpdateArtistRequest) (*UpdateArtistResponse, error)
	// *
	// Deletes an artist from our system.
	DeleteArtist(context.Context, *DeleteArtistRequest) (*DeleteArtistResponse, error)
	// *
	// Finds and retrieves artists matching the particular criteria.
	ListArtists(context.Context, *ListArtistsRequest) (*ListArtistsResponse, error)
	mustEmbedUnimplementedArtistServiceServer()
}

// UnimplementedArtistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArtistServiceServer struct {
}

func (UnimplementedArtistServiceServer) CreateArtist(context.Context, *CreateArtistRequest) (*CreateArtistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArtist not implemented")
}
func (UnimplementedArtistServiceServer) GetArtists(context.Context, *GetArtistsRequest) (*GetArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtists not implemented")
}
func (UnimplementedArtistServiceServer) UpdateArtist(context.Context, *UpdateArtistRequest) (*UpdateArtistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArtist not implemented")
}
func (UnimplementedArtistServiceServer) DeleteArtist(context.Context, *DeleteArtistRequest) (*DeleteArtistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtist not implemented")
}
func (UnimplementedArtistServiceServer) ListArtists(context.Context, *ListArtistsRequest) (*ListArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtists not implemented")
}
func (UnimplementedArtistServiceServer) mustEmbedUnimplementedArtistServiceServer() {}

// UnsafeArtistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtistServiceServer will
// result in compilation errors.
type UnsafeArtistServiceServer interface {
	mustEmbedUnimplementedArtistServiceServer()
}

func RegisterArtistServiceServer(s grpc.ServiceRegistrar, srv ArtistServiceServer) {
	s.RegisterService(&ArtistService_ServiceDesc, srv)
}

func _ArtistService_CreateArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).CreateArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_CreateArtist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).CreateArtist(ctx, req.(*CreateArtistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_GetArtists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).GetArtists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_GetArtists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).GetArtists(ctx, req.(*GetArtistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_UpdateArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).UpdateArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_UpdateArtist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).UpdateArtist(ctx, req.(*UpdateArtistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_DeleteArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).DeleteArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_DeleteArtist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).DeleteArtist(ctx, req.(*DeleteArtistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_ListArtists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).ListArtists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_ListArtists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).ListArtists(ctx, req.(*ListArtistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtistService_ServiceDesc is the grpc.ServiceDesc for ArtistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "musicservice.ArtistService",
	HandlerType: (*ArtistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArtist",
			Handler:    _ArtistService_CreateArtist_Handler,
		},
		{
			MethodName: "GetArtists",
			Handler:    _ArtistService_GetArtists_Handler,
		},
		{
			MethodName: "UpdateArtist",
			Handler:    _ArtistService_UpdateArtist_Handler,
		},
		{
			MethodName: "DeleteArtist",
			Handler:    _ArtistService_DeleteArtist_Handler,
		},
		{
			MethodName: "ListArtists",
			Handler:    _ArtistService_ListArtists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "musicservice/v1/artists.proto",
}
