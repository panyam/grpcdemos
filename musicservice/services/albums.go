package services

import (
	"context"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
	"log"
	protos "music.com/musicservice/gen/go/musicservice/v1"
	"strings"
)

type AlbumService struct {
	protos.UnimplementedAlbumServiceServer
	*EntityStore[protos.Album]
}

func NewAlbumService(estore *EntityStore[protos.Album]) *AlbumService {
	if estore == nil {
		estore = NewEntityStore[protos.Album]()
	}
	estore.IDSetter = func(album *protos.Album, id string) { album.Id = id }
	estore.IDGetter = func(album *protos.Album) string { return album.Id }

	estore.CreatedAtSetter = func(album *protos.Album, val *tspb.Timestamp) { album.CreatedAt = val }
	estore.CreatedAtGetter = func(album *protos.Album) *tspb.Timestamp { return album.CreatedAt }

	estore.UpdatedAtSetter = func(album *protos.Album, val *tspb.Timestamp) { album.UpdatedAt = val }
	estore.UpdatedAtGetter = func(album *protos.Album) *tspb.Timestamp { return album.UpdatedAt }

	return &AlbumService{
		EntityStore: estore,
	}
}

// Create a new Album
func (s *AlbumService) CreateAlbum(ctx context.Context, req *protos.CreateAlbumRequest) (resp *protos.CreateAlbumResponse, err error) {
	resp = &protos.CreateAlbumResponse{}
	resp.Album = s.EntityStore.Create(req.Album)
	return
}

// Batch gets multiple albums.
func (s *AlbumService) GetAlbums(ctx context.Context, req *protos.GetAlbumsRequest) (resp *protos.GetAlbumsResponse, err error) {
	log.Println("BatchGet for IDs: ", req.Ids)
	resp = &protos.GetAlbumsResponse{
		Albums: s.EntityStore.BatchGet(req.Ids),
	}
	return
}

// Updates specific fields of an Album
func (s *AlbumService) UpdateAlbum(ctx context.Context, req *protos.UpdateAlbumRequest) (resp *protos.UpdateAlbumResponse, err error) {
	resp = &protos.UpdateAlbumResponse{
		Album: s.EntityStore.Update(req.Album),
	}
	return
}

// Deletes an album from our system.
func (s *AlbumService) DeleteAlbum(ctx context.Context, req *protos.DeleteAlbumRequest) (resp *protos.DeleteAlbumResponse, err error) {
	resp = &protos.DeleteAlbumResponse{}
	s.EntityStore.Delete(req.Id)
	return
}

// Finds and retrieves albums matching the particular criteria.
func (s *AlbumService) ListAlbums(ctx context.Context, req *protos.ListAlbumsRequest) (resp *protos.ListAlbumsResponse, err error) {
	results := s.EntityStore.List(func(s1, s2 *protos.Album) bool {
		return strings.Compare(s1.Name, s2.Name) < 0
	})
	log.Println("Found Albums: ", results)
	resp = &protos.ListAlbumsResponse{Albums: results}
	return
}
