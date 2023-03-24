package services

import (
	"context"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
	"log"
	protos "music.com/musicservice/gen/go/musicservice/v1"
	"strings"
)

type LabelService struct {
	protos.UnimplementedLabelServiceServer
	*EntityStore[protos.Label]
}

func NewLabelService(estore *EntityStore[protos.Label]) *LabelService {
	if estore == nil {
		estore = NewEntityStore[protos.Label]()
	}
	estore.IDSetter = func(label *protos.Label, id string) { label.Id = id }
	estore.IDGetter = func(label *protos.Label) string { return label.Id }

	estore.CreatedAtSetter = func(label *protos.Label, val *tspb.Timestamp) { label.CreatedAt = val }
	estore.CreatedAtGetter = func(label *protos.Label) *tspb.Timestamp { return label.CreatedAt }

	estore.UpdatedAtSetter = func(label *protos.Label, val *tspb.Timestamp) { label.UpdatedAt = val }
	estore.UpdatedAtGetter = func(label *protos.Label) *tspb.Timestamp { return label.UpdatedAt }

	return &LabelService{
		EntityStore: estore,
	}
}

// Create a new Label
func (s *LabelService) CreateLabel(ctx context.Context, req *protos.CreateLabelRequest) (resp *protos.CreateLabelResponse, err error) {
	resp = &protos.CreateLabelResponse{}
	resp.Label = s.EntityStore.Create(req.Label)
	return
}

// Batch gets multiple labels.
func (s *LabelService) GetLabels(ctx context.Context, req *protos.GetLabelsRequest) (resp *protos.GetLabelsResponse, err error) {
	resp = &protos.GetLabelsResponse{
		Labels: s.EntityStore.BatchGet(req.Ids),
	}
	return
}

// Updates specific fields of an Label
func (s *LabelService) UpdateLabel(ctx context.Context, req *protos.UpdateLabelRequest) (resp *protos.UpdateLabelResponse, err error) {
	resp = &protos.UpdateLabelResponse{
		Label: s.EntityStore.Update(req.Label),
	}
	return
}

// Deletes an label from our system.
func (s *LabelService) DeleteLabel(ctx context.Context, req *protos.DeleteLabelRequest) (resp *protos.DeleteLabelResponse, err error) {
	resp = &protos.DeleteLabelResponse{}
	s.EntityStore.Delete(req.Id)
	return
}

// Finds and retrieves labels matching the particular criteria.
func (s *LabelService) ListLabels(ctx context.Context, req *protos.ListLabelsRequest) (resp *protos.ListLabelsResponse, err error) {
	results := s.EntityStore.List(func(s1, s2 *protos.Label) bool {
		return strings.Compare(s1.Name, s2.Name) < 0
	})
	log.Println("Found Labels: ", results)
	resp = &protos.ListLabelsResponse{Labels: results}
	return
}
