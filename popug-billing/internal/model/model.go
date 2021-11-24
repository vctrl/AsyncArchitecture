package model

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vctrl/async-architecture/popug-billing/internal/db"
)

type Model struct {
	transactions db.TransactionRepo
	tasks        db.TaskRepo
	users        db.UserRepoSQL

	producer *kafka.Producer
}

func (m *Model) CreatePlusTransaction(ctx context.Context, userID string, taskID string, amount int64) error {
	task, err := m.GetTask(ctx, taskID)
	if err != nil {
		return err
	}

	pid, id, err := m.transactions.Create(context, tran)
	if err != nil {
		return err
	}

	return pid, id, nil
}

func (m *Model) CreateMinusTransaction(ctx context.Context, taskID, userID string) (publicID string, id string, err error) {
	task, err := m.GetTask(ctx, taskID)
	if err != nil {
		return "", "", err
	}

	pid, id, err := m.transactions.Create(context, tran)
	if err != nil {
		return err
	}

	return pid, id, nil
}

func (m *Model) GetTask(ctx context.Context, id string) (*db.Task, error) {
	task, err := m.tasks.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
