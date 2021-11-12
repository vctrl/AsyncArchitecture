package api

import (
	"context"

	"github.com/vctrl/async-architecture/week_2/popug-tasks/internal/model"
	"github.com/vctrl/async-architecture/week_2/schema/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	tasks.UnimplementedTasksServer
	mdl *model.Model
}

func (s *server) Create(context.Context, *tasks.CreateRequest) (*tasks.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (s *server) AssignTo(context.Context, *tasks.AssignToRequest) (*tasks.AssignToResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignTo not implemented")
}
func (s *server) CreateAndAssignTo(context.Context, *tasks.CreateAndAssignToRequest) (*tasks.CreateAndAssignToResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAndAssignTo not implemented")
}
func (s *server) MarkAsDone(context.Context, *tasks.MarkAsDoneRequest) (*tasks.MarkAsDoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsDone not implemented")
}
func (s *server) Shuffle(context.Context, *tasks.ShuffleRequest) (*tasks.ShuffleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shuffle not implemented")
}
func (s *server) GetById(context.Context, *tasks.GetByIdRequest) (*tasks.GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (s *server) Update(context.Context, *tasks.UpdateRequest) (*tasks.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s *server) Delete(context.Context, *tasks.DeleteRequest) (*tasks.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func NewServer() *server {
	return &server{}
}
