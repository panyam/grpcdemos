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

func TestCreateSongAndFetch(t *testing.T) {
	svc := NewSongService(nil)
	RunTest(t,
		func(server *grpc.Server) {
			protos.RegisterSongServiceServer(server, svc)
		},
		func(ctx context.Context, conn *grpc.ClientConn) {
			client := protos.NewSongServiceClient(conn)
			resp, err := client.CreateSong(ctx, &protos.CreateSongRequest{
				Song: &protos.Song{
					Name:         "First Song",
					PerformerIds: []string{"1", "2", "3"},
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Song.Id, "1")
			assert.Equal(t, resp.Song.Name, "First Song")
			assert.Equal(t, resp.Song.PerformerIds, []string{"1", "2", "3"})

			// Create another
			resp, err = client.CreateSong(ctx, &protos.CreateSongRequest{
				Song: &protos.Song{
					Name:         "An awesome second song",
					PerformerIds: []string{"4", "2", "5"},
					Homepage:     "http://music.com/song/2",
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Song.Id, "2")
			assert.Equal(t, resp.Song.Name, "An awesome second song")
			assert.Equal(t, resp.Song.PerformerIds, []string{"4", "2", "5"})
			assert.Equal(t, resp.Song.Homepage, "http://music.com/song/2")

			resp2, err := client.GetSongs(ctx, &protos.GetSongsRequest{
				Ids: []string{"1", "2"},
			})
			assert.Equal(t, len(resp2.Songs), 2)
		})
}
