package model

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	uuid "github.com/satori/go.uuid"
	"github.com/vctrl/async-architecture/popug-billing/internal/db"
)

type Model struct {
	Transactions db.TransactionRepo
	Tasks        db.TaskRepo
	Users        db.UserRepo

	Producer *kafka.Producer
}

func (m *Model) CreatePlusTransaction(ctx context.Context, taskID, userID string) (publicID string, id string, err error) {
	task, err := m.GetTask(ctx, taskID)
	if err != nil {
		return "", "", err
	}

	pid, id, err := m.Transactions.Create(ctx, &db.Transaction{
		ID:       uuid.NewV4().String(),
		PublicID: uuid.NewV4().String(),
		UserID:   userID,
		TaskID:   taskID,
		Amount:   task.Prize,
	})
	if err != nil {
		return "", "", err
	}

	return pid, id, nil
}

// look what entities are used in production, what about entities on layers?
func (m *Model) CreateMinusTransaction(ctx context.Context, taskID, userID string) (publicID string, id string, err error) {
	task, err := m.GetTask(ctx, taskID)
	if err != nil {
		return "", "", err
	}

	pid, id, err := m.Transactions.Create(ctx, &db.Transaction{
		ID:       uuid.NewV4().String(),
		PublicID: uuid.NewV4().String(),
		UserID:   userID,
		TaskID:   taskID,
		Amount:   task.Cost,
	})
	if err != nil {
		return "", "", err
	}

	return pid, id, nil
}

func (m *Model) GetTask(ctx context.Context, id string) (*db.Task, error) {
	task, err := m.Tasks.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// todo
func (m *Model) CloseBillingCycle(ctx context.Context) error {
	// todo set status of transactions as closed and create new transaction to withdraw all money
	// don't create transaction, if balance is < 0
	return nil
}
