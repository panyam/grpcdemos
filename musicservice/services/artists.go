package services

import (
	"context"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
	"log"
	protos "music.com/musicservice/gen/go/musicservice/v1"
	"strings"
)

type ArtistService struct {
	protos.UnimplementedArtistServiceServer
	*EntityStore[protos.Artist]
}

func NewArtistService(estore *EntityStore[protos.Artist]) *ArtistService {
	if estore == nil {
		estore = NewEntityStore[protos.Artist]()
	}
	estore.IDSetter = func(artist *protos.Artist, id string) { artist.Id = id }
	estore.IDGetter = func(artist *protos.Artist) string { return artist.Id }

	estore.CreatedAtSetter = func(artist *protos.Artist, val *tspb.Timestamp) { artist.CreatedAt = val }
	estore.CreatedAtGetter = func(artist *protos.Artist) *tspb.Timestamp { return artist.CreatedAt }

	estore.UpdatedAtSetter = func(artist *protos.Artist, val *tspb.Timestamp) { artist.UpdatedAt = val }
	estore.UpdatedAtGetter = func(artist *protos.Artist) *tspb.Timestamp { return artist.UpdatedAt }

	return &ArtistService{
		EntityStore: estore,
	}
}

// Create a new Artist
func (s *ArtistService) CreateArtist(ctx context.Context, req *protos.CreateArtistRequest) (resp *protos.CreateArtistResponse, err error) {
	resp = &protos.CreateArtistResponse{}
	resp.Artist = s.EntityStore.Create(req.Artist)
	return
}

// Batch gets multiple artists.
func (s *ArtistService) GetArtists(ctx context.Context, req *protos.GetArtistsRequest) (resp *protos.GetArtistsResponse, err error) {
	resp = &protos.GetArtistsResponse{
		Artists: s.EntityStore.BatchGet(req.Ids),
	}
	return
}

// Updates specific fields of an Artist
func (s *ArtistService) UpdateArtist(ctx context.Context, req *protos.UpdateArtistRequest) (resp *protos.UpdateArtistResponse, err error) {
	resp = &protos.UpdateArtistResponse{
		Artist: s.EntityStore.Update(req.Artist),
	}
	return
}

// Deletes an artist from our system.
func (s *ArtistService) DeleteArtist(ctx context.Context, req *protos.DeleteArtistRequest) (resp *protos.DeleteArtistResponse, err error) {
	resp = &protos.DeleteArtistResponse{}
	s.EntityStore.Delete(req.Id)
	return
}

// Finds and retrieves artists matching the particular criteria.
func (s *ArtistService) ListArtists(ctx context.Context, req *protos.ListArtistsRequest) (resp *protos.ListArtistsResponse, err error) {
	results := s.EntityStore.List(func(s1, s2 *protos.Artist) bool {
		return strings.Compare(s1.Name, s2.Name) < 0
	})
	log.Println("Found Artists: ", results)
	resp = &protos.ListArtistsResponse{Artists: results}
	return
}
