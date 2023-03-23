package services

import (
	"context"

	protos "music.com/musicservice/gen/musicservice/v1"
)

type AlbumServiceServer struct {
}

func NewAlbumServiceServer() *AlbumServiceServer {
	return &AlbumServiceServer{}
}

// Create a new Album
func (s *AlbumServiceServer) CreateAlbum(ctx context.Context, req *protos.CreateAlbumRequest) (resp *protos.CreateAlbumResponse, err error) {
	resp = &protos.CreateAlbumResponse{}
	return
}

// Batch gets multiple albums.
func (s *AlbumServiceServer) GetAlbums(ctx context.Context, req *protos.GetAlbumsRequest) (resp *protos.GetAlbumsResponse, err error) {
	resp = &protos.GetAlbumsResponse{}
	return
}

// Updates specific fields of an Album
func (s *AlbumServiceServer) UpdateAlbum(ctx context.Context, req *protos.UpdateAlbumRequest) (resp *protos.UpdateAlbumResponse, err error) {
	resp = &protos.UpdateAlbumResponse{}
	return
}

// Deletes an album from our system.
func (s *AlbumServiceServer) DeleteAlbum(ctx context.Context, req *protos.DeleteAlbumRequest) (resp *protos.DeleteAlbumResponse, err error) {
	resp = &protos.DeleteAlbumResponse{}
	return
}

// Finds and retrieves albums matching the particular criteria.
func (s *AlbumServiceServer) ListAlbums(ctx context.Context, req *protos.ListAlbumsRequest) (resp *protos.ListAlbumsResponse, err error) {
	resp = &protos.ListAlbumsResponse{}
	return
}
