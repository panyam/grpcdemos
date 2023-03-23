package services

import (
	"context"

	protos "music.com/musicservice/gen/musicservice/v1"
)

type SongServiceServer struct {
}

func NewSongServiceServer() *SongServiceServer {
	return &SongServiceServer{}
}

// Create a new Song
func (s *SongServiceServer) CreateSong(ctx context.Context, req *protos.CreateSongRequest) (resp *protos.CreateSongResponse, err error) {
	resp = &protos.CreateSongResponse{}
	return
}

// Batch gets multiple songs.
func (s *SongServiceServer) GetSongs(ctx context.Context, req *protos.GetSongsRequest) (resp *protos.GetSongsResponse, err error) {
	resp = &protos.GetSongsResponse{}
	return
}

// Updates specific fields of an Song
func (s *SongServiceServer) UpdateSong(ctx context.Context, req *protos.UpdateSongRequest) (resp *protos.UpdateSongResponse, err error) {
	resp = &protos.UpdateSongResponse{}
	return
}

// Deletes an song from our system.
func (s *SongServiceServer) DeleteSong(ctx context.Context, req *protos.DeleteSongRequest) (resp *protos.DeleteSongResponse, err error) {
	resp = &protos.DeleteSongResponse{}
	return
}

// Finds and retrieves songs matching the particular criteria.
func (s *SongServiceServer) ListSongs(ctx context.Context, req *protos.ListSongsRequest) (resp *protos.ListSongsResponse, err error) {
	resp = &protos.ListSongsResponse{}
	return
}
