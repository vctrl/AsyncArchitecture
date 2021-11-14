package api

import (
	"context"

	"github.com/vctrl/async-architecture/week_2/popug-tasks/internal/db"
	"github.com/vctrl/async-architecture/week_2/popug-tasks/internal/model"
	"github.com/vctrl/async-architecture/week_2/schema/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	mdl *model.Model
	tasks.UnimplementedTasksServer
}

func New(dsn string) tasks.TasksServer {
	// todo config
	taskRepo := db.NewTaskRepoSQL(dsn)
	userRepo := db.NewUserRepositorySQL(dsn)
	mdl := &model.Model{
		Tasks: taskRepo,
		Users: userRepo,
	}
	return &server{
		mdl: mdl,
	}
}

func (s *server) CreateAndAssignTo(ctx context.Context, req *tasks.CreateAndAssignToRequest) (*tasks.CreateAndAssignToResponse, error) {
	publicID, id, err := s.mdl.CreateAndAssignTo(ctx, req.TaskInfo.Description, req.AssignToId)
	if err != nil {
		return nil, err
	}

	return &tasks.CreateAndAssignToResponse{
		Status:   nil,
		PublicId: publicID,
		Id:       id,
	}, nil
}

func (s *server) MarkAsDone(ctx context.Context, req *tasks.MarkAsDoneRequest) (*tasks.MarkAsDoneResponse, error) {
	err := s.mdl.MarkAsDone(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}
	return &tasks.MarkAsDoneResponse{Status: &tasks.Status{
		Code: int32(codes.OK),
		Msg:  "",
	}}, nil
}

func (s *server) Shuffle(ctx context.Context, req *tasks.ShuffleRequest) (*tasks.ShuffleResponse, error) {
	err := s.mdl.Shuffle(ctx)
	if err != nil {
		return nil, err
	}

	return &tasks.ShuffleResponse{
		Status: &tasks.Status{
			Code: int32(codes.OK),
			Msg:  "",
		},
	}, nil
}

func (s *server) GetAll(ctx context.Context, req *tasks.GetAllRequest) (*tasks.GetAllResponse, error) {
	ts, err := s.mdl.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*tasks.TaskInfo, 0, len(ts))
	for _, t := range ts {
		res = append(res, &tasks.TaskInfo{
			Description: t.Description,
			Done:        t.Done,
		})
	}

	return &tasks.GetAllResponse{
		Status: &tasks.Status{
			Code: int32(codes.OK),
			Msg:  "",
		},
		Tasks: res,
	}, nil
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
