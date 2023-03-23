package services

import (
	"context"

	protos "music.com/musicservice/gen/musicservice/v1"
)

type ArtistServiceServer struct {
}

func NewArtistServiceServer() *ArtistServiceServer {
	return &ArtistServiceServer{}
}

// Create a new Artist
func (s *ArtistServiceServer) CreateArtist(ctx context.Context, req *protos.CreateArtistRequest) (resp *protos.CreateArtistResponse, err error) {
	resp = &protos.CreateArtistResponse{}
	return
}

// Batch gets multiple artists.
func (s *ArtistServiceServer) GetArtists(ctx context.Context, req *protos.GetArtistsRequest) (resp *protos.GetArtistsResponse, err error) {
	resp = &protos.GetArtistsResponse{}
	return
}

// Updates specific fields of an Artist
func (s *ArtistServiceServer) UpdateArtist(ctx context.Context, req *protos.UpdateArtistRequest) (resp *protos.UpdateArtistResponse, err error) {
	resp = &protos.UpdateArtistResponse{}
	return
}

// Deletes an artist from our system.
func (s *ArtistServiceServer) DeleteArtist(ctx context.Context, req *protos.DeleteArtistRequest) (resp *protos.DeleteArtistResponse, err error) {
	resp = &protos.DeleteArtistResponse{}
	return
}

// Finds and retrieves artists matching the particular criteria.
func (s *ArtistServiceServer) ListArtists(ctx context.Context, req *protos.ListArtistsRequest) (resp *protos.ListArtistsResponse, err error) {
	resp = &protos.ListArtistsResponse{}
	return
}
