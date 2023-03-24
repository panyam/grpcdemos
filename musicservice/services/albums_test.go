package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	protos "music.com/musicservice/gen/go/musicservice/v1"
	"testing"
)

func TestCreateAlbumAndFetch(t *testing.T) {
	svc := NewAlbumService(nil)
	RunTest(t,
		func(server *grpc.Server) {
			protos.RegisterAlbumServiceServer(server, svc)
		},
		func(ctx context.Context, conn *grpc.ClientConn) {
			client := protos.NewAlbumServiceClient(conn)
			resp, err := client.CreateAlbum(ctx, &protos.CreateAlbumRequest{
				Album: &protos.Album{
					Name:    "First Album",
					SongIds: []string{"1", "2", "3"},
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Album.Id, "1")
			assert.Equal(t, resp.Album.Name, "First Album")
			assert.Equal(t, resp.Album.SongIds, []string{"1", "2", "3"})

			// Create another
			resp, err = client.CreateAlbum(ctx, &protos.CreateAlbumRequest{
				Album: &protos.Album{
					Name:    "An awesome second album",
					SongIds: []string{"4", "2", "5"},
					Venue:   "Australia",
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Album.Id, "2")
			assert.Equal(t, resp.Album.Name, "An awesome second album")
			assert.Equal(t, resp.Album.SongIds, []string{"4", "2", "5"})
			assert.Equal(t, resp.Album.Venue, "Australia")

			resp2, err := client.GetAlbums(ctx, &protos.GetAlbumsRequest{
				Ids: []string{"1", "2"},
			})
			assert.Equal(t, len(resp2.Albums), 2)
		})
}
