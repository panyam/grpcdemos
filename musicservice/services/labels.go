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

type LabelServiceServer struct {
	protos.UnimplementedLabelServiceServer
	Entities map[string]*protos.Label
	IDCount  int
}

func NewLabelServiceServer() *LabelServiceServer {
	return &LabelServiceServer{
		Entities: make(map[string]*protos.Label),
	}
}

// Create a new Label
func (s *LabelServiceServer) CreateLabel(ctx context.Context, req *protos.CreateLabelRequest) (resp *protos.CreateLabelResponse, err error) {
	resp = &protos.CreateLabelResponse{}
	resp.Label = req.Label
	s.IDCount++
	req.Label.Id = fmt.Sprintf("%d", s.IDCount)
	req.Label.CreatedAt = tspb.New(time.Now())
	req.Label.UpdatedAt = tspb.New(time.Now())
	s.Entities[req.Label.Id] = req.Label

	// Ideally we want:
	// resp.Label = s.BaseEntityService.CreateEntity(req.Label)
	return
}

// Batch gets multiple labels.
func (s *LabelServiceServer) GetLabels(ctx context.Context, req *protos.GetLabelsRequest) (resp *protos.GetLabelsResponse, err error) {
	resp = &protos.GetLabelsResponse{
		Labels: make(map[string]*protos.Label),
	}
	for _, labelid := range req.Ids {
		if label, ok := s.Entities[labelid]; ok {
			resp.Labels[labelid] = label
		}
	}
	return
}

// Updates specific fields of an Label
func (s *LabelServiceServer) UpdateLabel(ctx context.Context, req *protos.UpdateLabelRequest) (resp *protos.UpdateLabelResponse, err error) {
	resp = &protos.UpdateLabelResponse{
		Label: req.Label,
	}
	resp.Label.UpdatedAt = tspb.New(time.Now())
	return
}

// Deletes an label from our system.
func (s *LabelServiceServer) DeleteLabel(ctx context.Context, req *protos.DeleteLabelRequest) (resp *protos.DeleteLabelResponse, err error) {
	resp = &protos.DeleteLabelResponse{}
	if _, ok := s.Entities[req.Id]; ok {
		delete(s.Entities, req.Id)
	}
	return
}

// Finds and retrieves labels matching the particular criteria.
func (s *LabelServiceServer) ListLabels(ctx context.Context, req *protos.ListLabelsRequest) (resp *protos.ListLabelsResponse, err error) {
	resp = &protos.ListLabelsResponse{}
	for _, label := range s.Entities {
		resp.Labels = append(resp.Labels, label)
	}
	// Sort in reverse order of name
	sort.Slice(resp.Labels, func(idx1, idx2 int) bool {
		label1 := resp.Labels[idx1]
		label2 := resp.Labels[idx2]
		return strings.Compare(label1.Name, label2.Name) < 0
	})
	return
}
