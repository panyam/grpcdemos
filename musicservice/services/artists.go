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

type ArtistServiceServer struct {
	protos.UnimplementedArtistServiceServer
	Entities map[string]*protos.Artist
	IDCount  int
}

func NewArtistServiceServer() *ArtistServiceServer {
	return &ArtistServiceServer{
		Entities: make(map[string]*protos.Artist),
	}
}

// Create a new Artist
func (s *ArtistServiceServer) CreateArtist(ctx context.Context, req *protos.CreateArtistRequest) (resp *protos.CreateArtistResponse, err error) {
	resp = &protos.CreateArtistResponse{}
	resp.Artist = req.Artist
	s.IDCount++
	req.Artist.Id = fmt.Sprintf("%d", s.IDCount)
	req.Artist.CreatedAt = tspb.New(time.Now())
	req.Artist.UpdatedAt = tspb.New(time.Now())
	s.Entities[req.Artist.Id] = req.Artist

	// Ideally we want:
	// resp.Artist = s.BaseEntityService.CreateEntity(req.Artist)
	return
}

// Batch gets multiple artists.
func (s *ArtistServiceServer) GetArtists(ctx context.Context, req *protos.GetArtistsRequest) (resp *protos.GetArtistsResponse, err error) {
	resp = &protos.GetArtistsResponse{
		Artists: make(map[string]*protos.Artist),
	}
	for _, artistid := range req.Ids {
		if artist, ok := s.Entities[artistid]; ok {
			resp.Artists[artistid] = artist
		}
	}
	return
}

// Updates specific fields of an Artist
func (s *ArtistServiceServer) UpdateArtist(ctx context.Context, req *protos.UpdateArtistRequest) (resp *protos.UpdateArtistResponse, err error) {
	resp = &protos.UpdateArtistResponse{
		Artist: req.Artist,
	}
	resp.Artist.UpdatedAt = tspb.New(time.Now())
	return
}

// Deletes an artist from our system.
func (s *ArtistServiceServer) DeleteArtist(ctx context.Context, req *protos.DeleteArtistRequest) (resp *protos.DeleteArtistResponse, err error) {
	resp = &protos.DeleteArtistResponse{}
	if _, ok := s.Entities[req.Id]; ok {
		delete(s.Entities, req.Id)
	}
	return
}

// Finds and retrieves artists matching the particular criteria.
func (s *ArtistServiceServer) ListArtists(ctx context.Context, req *protos.ListArtistsRequest) (resp *protos.ListArtistsResponse, err error) {
	resp = &protos.ListArtistsResponse{}
	for _, artist := range s.Entities {
		resp.Artists = append(resp.Artists, artist)
	}
	// Sort in reverse order of name
	sort.Slice(resp.Artists, func(idx1, idx2 int) bool {
		artist1 := resp.Artists[idx1]
		artist2 := resp.Artists[idx2]
		return strings.Compare(artist1.Name, artist2.Name) < 0
	})
	return
}
