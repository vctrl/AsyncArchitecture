package api

import (
	"context"
	"github.com/vctrl/async-architecture/popug-analytics/internal/model"
	"github.com/vctrl/async-architecture/schema/analytics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	mdl *model.Model
	analytics.UnimplementedAnalyticsServer
}

func (s *server) GetTodayEarnings(context.Context, *analytics.GetTodayEarningsRequest) (*analytics.GetTodayEarningsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodayRevenue not implemented")
}
func (s *server) GetNonProfitPopugCount(context.Context, *analytics.GetNonProfitPopugCountRequest) (*analytics.GetNonProfitPopugCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnprofitableCount not implemented")
}
func (s *server) GetMostExpensiveTask(context.Context, *analytics.GetMostExpensiveTaskRequest) (*analytics.GetMostExpensiveTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMostExpensiveTask not implemented")
}