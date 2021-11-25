package api

import (
	"context"
	"github.com/vctrl/async-architecture/popug-billing/internal/db"
	"github.com/vctrl/async-architecture/popug-billing/internal/model"
	"github.com/vctrl/async-architecture/schema/billing"
	"google.golang.org/grpc/codes"
)

type server struct {
	Mdl *model.Model
	billing.UnimplementedBillingServer
}

func New(dsn string) *server {
	tasks := db.NewTaskRepoSQL(dsn)
	users := db.NewUserRepositorySQL(dsn)
	transactions := db.NewTransactionRepoSQL(dsn)
	mdl := &model.Model{
		Transactions: transactions,
		Tasks:        tasks,
		Users:        users,
		Producer:     nil,
	}
	return &server{
		Mdl: mdl,
	}
}

func (s *server) CreatePlusTransaction(ctx context.Context, req *billing.CreatePlusTransactionRequest) (*billing.CreatePlusTransactionResponse, error) {
	return &billing.CreatePlusTransactionResponse{Status: &billing.Status{Code: int32(codes.OK)}}, nil
}

func (s *server) CreateMinusTransaction(ctx context.Context, req *billing.CreateMinusTransactionRequest) (*billing.CreateMinusTransactionResponse, error) {
	return &billing.CreateMinusTransactionResponse{Status: &billing.Status{Code: int32(codes.OK)}}, nil
}

func (s *server) CloseBillingCycle(ctx context.Context, req *billing.CloseBillingCycleRequest) (*billing.CloseBillingCycleResponse, error) {
	return &billing.CloseBillingCycleResponse{Status: &billing.Status{Code: int32(codes.OK)}}, nil
}
