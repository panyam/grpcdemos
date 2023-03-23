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

type AlbumServiceServer struct {
	protos.UnimplementedAlbumServiceServer
	Entities map[string]*protos.Album
	IDCount  int
}

func NewAlbumServiceServer() *AlbumServiceServer {
	return &AlbumServiceServer{
		Entities: make(map[string]*protos.Album),
	}
}

// Create a new Album
func (s *AlbumServiceServer) CreateAlbum(ctx context.Context, req *protos.CreateAlbumRequest) (resp *protos.CreateAlbumResponse, err error) {
	resp = &protos.CreateAlbumResponse{}
	resp.Album = req.Album
	s.IDCount++
	req.Album.Id = fmt.Sprintf("%d", s.IDCount)
	req.Album.CreatedAt = tspb.New(time.Now())
	req.Album.UpdatedAt = tspb.New(time.Now())
	s.Entities[req.Album.Id] = req.Album

	// Ideally we want:
	// resp.Album = s.BaseEntityService.CreateEntity(req.Album)
	return
}

// Batch gets multiple albums.
func (s *AlbumServiceServer) GetAlbums(ctx context.Context, req *protos.GetAlbumsRequest) (resp *protos.GetAlbumsResponse, err error) {
	resp = &protos.GetAlbumsResponse{
		Albums: make(map[string]*protos.Album),
	}
	for _, albumid := range req.Ids {
		if album, ok := s.Entities[albumid]; ok {
			resp.Albums[albumid] = album
		}
	}
	return
}

// Updates specific fields of an Album
func (s *AlbumServiceServer) UpdateAlbum(ctx context.Context, req *protos.UpdateAlbumRequest) (resp *protos.UpdateAlbumResponse, err error) {
	resp = &protos.UpdateAlbumResponse{
		Album: req.Album,
	}
	resp.Album.UpdatedAt = tspb.New(time.Now())
	return
}

// Deletes an album from our system.
func (s *AlbumServiceServer) DeleteAlbum(ctx context.Context, req *protos.DeleteAlbumRequest) (resp *protos.DeleteAlbumResponse, err error) {
	resp = &protos.DeleteAlbumResponse{}
	if _, ok := s.Entities[req.Id]; ok {
		delete(s.Entities, req.Id)
	}
	return
}

// Finds and retrieves albums matching the particular criteria.
func (s *AlbumServiceServer) ListAlbums(ctx context.Context, req *protos.ListAlbumsRequest) (resp *protos.ListAlbumsResponse, err error) {
	resp = &protos.ListAlbumsResponse{}
	for _, album := range s.Entities {
		resp.Albums = append(resp.Albums, album)
	}
	// Sort in reverse order of name
	sort.Slice(resp.Albums, func(idx1, idx2 int) bool {
		album1 := resp.Albums[idx1]
		album2 := resp.Albums[idx2]
		return strings.Compare(album1.Name, album2.Name) < 0
	})
	return
}
