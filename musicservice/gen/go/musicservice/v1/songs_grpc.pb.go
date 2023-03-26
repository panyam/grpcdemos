// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: songs.proto

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
	SongService_CreateSong_FullMethodName = "/musicservice.SongService/CreateSong"
	SongService_GetSongs_FullMethodName   = "/musicservice.SongService/GetSongs"
	SongService_UpdateSong_FullMethodName = "/musicservice.SongService/UpdateSong"
	SongService_DeleteSong_FullMethodName = "/musicservice.SongService/DeleteSong"
	SongService_ListSongs_FullMethodName  = "/musicservice.SongService/ListSongs"
)

// SongServiceClient is the client API for SongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SongServiceClient interface {
	// *
	// Create a new Song
	CreateSong(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*CreateSongResponse, error)
	// *
	// Batch gets multiple songs.
	GetSongs(ctx context.Context, in *GetSongsRequest, opts ...grpc.CallOption) (*GetSongsResponse, error)
	// *
	// Updates specific fields of an Song
	UpdateSong(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*UpdateSongResponse, error)
	// *
	// Deletes an song from our system.
	DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*DeleteSongResponse, error)
	// *
	// Finds and retrieves songs matching the particular criteria.
	ListSongs(ctx context.Context, in *ListSongsRequest, opts ...grpc.CallOption) (*ListSongsResponse, error)
}

type songServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSongServiceClient(cc grpc.ClientConnInterface) SongServiceClient {
	return &songServiceClient{cc}
}

func (c *songServiceClient) CreateSong(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*CreateSongResponse, error) {
	out := new(CreateSongResponse)
	err := c.cc.Invoke(ctx, SongService_CreateSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) GetSongs(ctx context.Context, in *GetSongsRequest, opts ...grpc.CallOption) (*GetSongsResponse, error) {
	out := new(GetSongsResponse)
	err := c.cc.Invoke(ctx, SongService_GetSongs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) UpdateSong(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*UpdateSongResponse, error) {
	out := new(UpdateSongResponse)
	err := c.cc.Invoke(ctx, SongService_UpdateSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*DeleteSongResponse, error) {
	out := new(DeleteSongResponse)
	err := c.cc.Invoke(ctx, SongService_DeleteSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) ListSongs(ctx context.Context, in *ListSongsRequest, opts ...grpc.CallOption) (*ListSongsResponse, error) {
	out := new(ListSongsResponse)
	err := c.cc.Invoke(ctx, SongService_ListSongs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SongServiceServer is the server API for SongService service.
// All implementations must embed UnimplementedSongServiceServer
// for forward compatibility
type SongServiceServer interface {
	// *
	// Create a new Song
	CreateSong(context.Context, *CreateSongRequest) (*CreateSongResponse, error)
	// *
	// Batch gets multiple songs.
	GetSongs(context.Context, *GetSongsRequest) (*GetSongsResponse, error)
	// *
	// Updates specific fields of an Song
	UpdateSong(context.Context, *UpdateSongRequest) (*UpdateSongResponse, error)
	// *
	// Deletes an song from our system.
	DeleteSong(context.Context, *DeleteSongRequest) (*DeleteSongResponse, error)
	// *
	// Finds and retrieves songs matching the particular criteria.
	ListSongs(context.Context, *ListSongsRequest) (*ListSongsResponse, error)
	mustEmbedUnimplementedSongServiceServer()
}

// UnimplementedSongServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSongServiceServer struct {
}

func (UnimplementedSongServiceServer) CreateSong(context.Context, *CreateSongRequest) (*CreateSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSong not implemented")
}
func (UnimplementedSongServiceServer) GetSongs(context.Context, *GetSongsRequest) (*GetSongsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSongs not implemented")
}
func (UnimplementedSongServiceServer) UpdateSong(context.Context, *UpdateSongRequest) (*UpdateSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSong not implemented")
}
func (UnimplementedSongServiceServer) DeleteSong(context.Context, *DeleteSongRequest) (*DeleteSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedSongServiceServer) ListSongs(context.Context, *ListSongsRequest) (*ListSongsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSongs not implemented")
}
func (UnimplementedSongServiceServer) mustEmbedUnimplementedSongServiceServer() {}

// UnsafeSongServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SongServiceServer will
// result in compilation errors.
type UnsafeSongServiceServer interface {
	mustEmbedUnimplementedSongServiceServer()
}

func RegisterSongServiceServer(s grpc.ServiceRegistrar, srv SongServiceServer) {
	s.RegisterService(&SongService_ServiceDesc, srv)
}

func _SongService_CreateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).CreateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SongService_CreateSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).CreateSong(ctx, req.(*CreateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_GetSongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSongsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).GetSongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SongService_GetSongs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).GetSongs(ctx, req.(*GetSongsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_UpdateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).UpdateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SongService_UpdateSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).UpdateSong(ctx, req.(*UpdateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SongService_DeleteSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).DeleteSong(ctx, req.(*DeleteSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_ListSongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSongsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).ListSongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SongService_ListSongs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).ListSongs(ctx, req.(*ListSongsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SongService_ServiceDesc is the grpc.ServiceDesc for SongService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SongService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "musicservice.SongService",
	HandlerType: (*SongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSong",
			Handler:    _SongService_CreateSong_Handler,
		},
		{
			MethodName: "GetSongs",
			Handler:    _SongService_GetSongs_Handler,
		},
		{
			MethodName: "UpdateSong",
			Handler:    _SongService_UpdateSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _SongService_DeleteSong_Handler,
		},
		{
			MethodName: "ListSongs",
			Handler:    _SongService_ListSongs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "songs.proto",
}
