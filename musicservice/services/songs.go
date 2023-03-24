package services

import (
	"context"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
	"log"
	protos "music.com/musicservice/gen/go/musicservice/v1"
	"strings"
)

type SongService struct {
	protos.UnimplementedSongServiceServer
	*EntityStore[protos.Song]
}

func NewSongService(estore *EntityStore[protos.Song]) *SongService {
	if estore == nil {
		estore = NewEntityStore[protos.Song]()
	}
	estore.IDSetter = func(song *protos.Song, id string) { song.Id = id }
	estore.IDGetter = func(song *protos.Song) string { return song.Id }

	estore.CreatedAtSetter = func(song *protos.Song, val *tspb.Timestamp) { song.CreatedAt = val }
	estore.CreatedAtGetter = func(song *protos.Song) *tspb.Timestamp { return song.CreatedAt }

	estore.UpdatedAtSetter = func(song *protos.Song, val *tspb.Timestamp) { song.UpdatedAt = val }
	estore.UpdatedAtGetter = func(song *protos.Song) *tspb.Timestamp { return song.UpdatedAt }

	return &SongService{
		EntityStore: estore,
	}
}

// Create a new Song
func (s *SongService) CreateSong(ctx context.Context, req *protos.CreateSongRequest) (resp *protos.CreateSongResponse, err error) {
	resp = &protos.CreateSongResponse{}
	resp.Song = s.EntityStore.Create(req.Song)
	return
}

// Batch gets multiple songs.
func (s *SongService) GetSongs(ctx context.Context, req *protos.GetSongsRequest) (resp *protos.GetSongsResponse, err error) {
	log.Println("BatchGet for IDs: ", req.Ids)
	resp = &protos.GetSongsResponse{
		Songs: s.EntityStore.BatchGet(req.Ids),
	}
	return
}

// Updates specific fields of an Song
func (s *SongService) UpdateSong(ctx context.Context, req *protos.UpdateSongRequest) (resp *protos.UpdateSongResponse, err error) {
	resp = &protos.UpdateSongResponse{
		Song: s.EntityStore.Update(req.Song),
	}
	return
}

// Deletes an song from our system.
func (s *SongService) DeleteSong(ctx context.Context, req *protos.DeleteSongRequest) (resp *protos.DeleteSongResponse, err error) {
	resp = &protos.DeleteSongResponse{}
	s.EntityStore.Delete(req.Id)
	return
}

// Finds and retrieves songs matching the particular criteria.
func (s *SongService) ListSongs(ctx context.Context, req *protos.ListSongsRequest) (resp *protos.ListSongsResponse, err error) {
	results := s.EntityStore.List(func(s1, s2 *protos.Song) bool {
		return strings.Compare(s1.Name, s2.Name) < 0
	})
	log.Println("Found Songs: ", results)
	resp = &protos.ListSongsResponse{Songs: results}
	return
}
