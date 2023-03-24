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

func TestCreateArtistAndFetch(t *testing.T) {
	svc := NewArtistService(nil)
	RunTest(t,
		func(server *grpc.Server) {
			protos.RegisterArtistServiceServer(server, svc)
		},
		func(ctx context.Context, conn *grpc.ClientConn) {
			client := protos.NewArtistServiceClient(conn)
			resp, err := client.CreateArtist(ctx, &protos.CreateArtistRequest{
				Artist: &protos.Artist{
					Name: "First Artist",
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Artist.Id, "1")
			assert.Equal(t, resp.Artist.Name, "First Artist")

			// Create another
			resp, err = client.CreateArtist(ctx, &protos.CreateArtistRequest{
				Artist: &protos.Artist{
					Name:    "An awesome second artist",
					Country: "Canada",
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Artist.Id, "2")
			assert.Equal(t, resp.Artist.Name, "An awesome second artist")
			assert.Equal(t, resp.Artist.Country, "Canada")

			resp2, err := client.GetArtists(ctx, &protos.GetArtistsRequest{
				Ids: []string{"1", "2"},
			})
			assert.Equal(t, len(resp2.Artists), 2)
		})
}
