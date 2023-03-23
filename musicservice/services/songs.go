package services

import (
	"context"
	protos "music.com/musicservice/gen/go/musicservice/v1"
	"strings"
)

type SongServiceServer struct {
	protos.UnimplementedSongServiceServer
	*EntityService[protos.Song]
}

func NewSongServiceServer() *SongServiceServer {
	return &SongServiceServer{
		EntityService: NewEntityService[protos.Song](),
	}
}

// Create a new Song
func (s *SongServiceServer) CreateSong(ctx context.Context, req *protos.CreateSongRequest) (resp *protos.CreateSongResponse, err error) {
	resp = &protos.CreateSongResponse{}
	resp.Song = s.EntityService.Create(req.Song)

	// Ideally we want:
	// resp.Song = s.BaseEntityService.CreateEntity(req.Song)
	return
}

// Batch gets multiple songs.
func (s *SongServiceServer) GetSongs(ctx context.Context, req *protos.GetSongsRequest) (resp *protos.GetSongsResponse, err error) {
	resp = &protos.GetSongsResponse{
		Songs: s.EntityService.BatchGet(req.Ids),
	}
	return
}

// Updates specific fields of an Song
func (s *SongServiceServer) UpdateSong(ctx context.Context, req *protos.UpdateSongRequest) (resp *protos.UpdateSongResponse, err error) {
	resp = &protos.UpdateSongResponse{
		Song: s.EntityService.Update(req.Song),
	}
	return
}

// Deletes an song from our system.
func (s *SongServiceServer) DeleteSong(ctx context.Context, req *protos.DeleteSongRequest) (resp *protos.DeleteSongResponse, err error) {
	resp = &protos.DeleteSongResponse{}
	s.EntityService.Delete(req.Id)
	return
}

// Finds and retrieves songs matching the particular criteria.
func (s *SongServiceServer) ListSongs(ctx context.Context, req *protos.ListSongsRequest) (resp *protos.ListSongsResponse, err error) {
	resp = &protos.ListSongsResponse{
		Songs: s.EntityService.List(func(s1, s2 *protos.Song) bool {
			return strings.Compare(s1.Name, s2.Name) < 0
		}),
	}
	return
}
