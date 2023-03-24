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

func TestCreateLabelAndFetch(t *testing.T) {
	svc := NewLabelService(nil)
	RunTest(t,
		func(server *grpc.Server) {
			protos.RegisterLabelServiceServer(server, svc)
		},
		func(ctx context.Context, conn *grpc.ClientConn) {
			client := protos.NewLabelServiceClient(conn)
			resp, err := client.CreateLabel(ctx, &protos.CreateLabelRequest{
				Label: &protos.Label{
					Name: "First Label",
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Label.Id, "1")
			assert.Equal(t, resp.Label.Name, "First Label")

			// Create another
			resp, err = client.CreateLabel(ctx, &protos.CreateLabelRequest{
				Label: &protos.Label{
					Name: "An awesome second label",
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Label.Id, "2")
			assert.Equal(t, resp.Label.Name, "An awesome second label")

			resp2, err := client.GetLabels(ctx, &protos.GetLabelsRequest{
				Ids: []string{"1", "2"},
			})
			assert.Equal(t, len(resp2.Labels), 2)
		})
}
