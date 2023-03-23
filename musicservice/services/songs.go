package services

import (
	"context"
	"fmt"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
	protos "music.com/musicservice/gen/go/musicservice/v1"
	"sort"
	"strings"
	"time"
)

type SongServiceServer struct {
	protos.UnimplementedSongServiceServer
	Entities map[string]*protos.Song
	IDCount  int
}

func NewSongServiceServer() *SongServiceServer {
	return &SongServiceServer{
		Entities: make(map[string]*protos.Song),
	}
}

// Create a new Song
func (s *SongServiceServer) CreateSong(ctx context.Context, req *protos.CreateSongRequest) (resp *protos.CreateSongResponse, err error) {
	resp = &protos.CreateSongResponse{}
	resp.Song = req.Song
	s.IDCount++
	req.Song.Id = fmt.Sprintf("%d", s.IDCount)
	req.Song.CreatedAt = tspb.New(time.Now())
	req.Song.UpdatedAt = tspb.New(time.Now())
	s.Entities[req.Song.Id] = req.Song

	// Ideally we want:
	// resp.Song = s.BaseEntityService.CreateEntity(req.Song)
	return
}

// Batch gets multiple songs.
func (s *SongServiceServer) GetSongs(ctx context.Context, req *protos.GetSongsRequest) (resp *protos.GetSongsResponse, err error) {
	resp = &protos.GetSongsResponse{
		Songs: make(map[string]*protos.Song),
	}
	for _, songid := range req.Ids {
		if song, ok := s.Entities[songid]; ok {
			resp.Songs[songid] = song
		}
	}
	return
}

// Updates specific fields of an Song
func (s *SongServiceServer) UpdateSong(ctx context.Context, req *protos.UpdateSongRequest) (resp *protos.UpdateSongResponse, err error) {
	resp = &protos.UpdateSongResponse{
		Song: req.Song,
	}
	resp.Song.UpdatedAt = tspb.New(time.Now())
	return
}

// Deletes an song from our system.
func (s *SongServiceServer) DeleteSong(ctx context.Context, req *protos.DeleteSongRequest) (resp *protos.DeleteSongResponse, err error) {
	resp = &protos.DeleteSongResponse{}
	if _, ok := s.Entities[req.Id]; ok {
		delete(s.Entities, req.Id)
	}
	return
}

// Finds and retrieves songs matching the particular criteria.
func (s *SongServiceServer) ListSongs(ctx context.Context, req *protos.ListSongsRequest) (resp *protos.ListSongsResponse, err error) {
	resp = &protos.ListSongsResponse{}
	for _, song := range s.Entities {
		resp.Songs = append(resp.Songs, song)
	}
	// Sort in reverse order of name
	sort.Slice(resp.Songs, func(idx1, idx2 int) bool {
		song1 := resp.Songs[idx1]
		song2 := resp.Songs[idx2]
		return strings.Compare(song1.Name, song2.Name) < 0
	})
	return
}
