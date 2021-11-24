package api

import (
	"context"
	"github.com/vctrl/async-architecture/popug-billing/internal/model"
	"github.com/vctrl/async-architecture/schema/billing"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	mdl *model.Model
	billing.UnimplementedBillingServer
}

func New() *server {
	return &server{}
}

func (s *server) CreatePlusTransaction(ctx context.Context, req *billing.CreateTransactionRequest) (*billing.CreateTransactionResponse, error) {
	s.mdl.CreatePlusTransaction(ctx)
	return &billing.CreateTransactionResponse{Status: &billing.Status{Code: int32(codes.OK)}}, nil
}

func (s *server) CreateMinusTransaction(ctx context.Context, req *billing.CreateTransactionRequest) (*billing.CreateTransactionResponse, error) {
	s.mdl.CreateMinusTransaction(ctx)
	return &billing.CreateTransactionResponse{Status: &billing.Status{Code: int32(codes.OK)}}, nil
}

func (s *server) CloseBillingCycle(ctx context.Context, req *billing.CloseBillingCycleRequest) (*billing.CloseBillingCycleResponse, error) {
	s.mdl.CloseBillingCycle(ctx)
	return &billing.CloseBillingCycleResponse{Status: &billing.Status{Code: int32(codes.OK)}}, nil
}

func (s *server) CreateTask() {}
