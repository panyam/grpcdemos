package services

import (
	"context"

	protos "music.com/musicservice/gen/musicservice/v1"
)

type LabelServiceServer struct {
}

func NewLabelServiceServer() *LabelServiceServer {
	return &LabelServiceServer{}
}

// Create a new Label
func (s *LabelServiceServer) CreateLabel(ctx context.Context, req *protos.CreateLabelRequest) (resp *protos.CreateLabelResponse, err error) {
	resp = &protos.CreateLabelResponse{}
	return
}

// Batch gets multiple labels.
func (s *LabelServiceServer) GetLabels(ctx context.Context, req *protos.GetLabelsRequest) (resp *protos.GetLabelsResponse, err error) {
	resp = &protos.GetLabelsResponse{}
	return
}

// Updates specific fields of an Label
func (s *LabelServiceServer) UpdateLabel(ctx context.Context, req *protos.UpdateLabelRequest) (resp *protos.UpdateLabelResponse, err error) {
	resp = &protos.UpdateLabelResponse{}
	return
}

// Deletes an label from our system.
func (s *LabelServiceServer) DeleteLabel(ctx context.Context, req *protos.DeleteLabelRequest) (resp *protos.DeleteLabelResponse, err error) {
	resp = &protos.DeleteLabelResponse{}
	return
}

// Finds and retrieves labels matching the particular criteria.
func (s *LabelServiceServer) ListLabels(ctx context.Context, req *protos.ListLabelsRequest) (resp *protos.ListLabelsResponse, err error) {
	resp = &protos.ListLabelsResponse{}
	return
}
