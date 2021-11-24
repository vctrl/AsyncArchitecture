package model

import (
	"context"
	"math/rand"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vctrl/async-architecture/popug-tasks/internal/db"
	tasks2 "github.com/vctrl/async-architecture/week_2/schema/tasks"
)

type Model struct {
	Tasks    db.TaskRepo
	Users    db.UserRepo
	Producer *kafka.Producer
}

type NewTaskAssignedTo struct {
	TaskID        string
	NewAssignedTo string
}

func (m *Model) Shuffle(ctx context.Context) ([]*NewTaskAssignedTo, error) {
	ids, err := m.Users.GetAllUserIDs(ctx)
	if err != nil {
		return nil, err
	}

	tasks, err := m.Tasks.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	changeInfo := make([]*NewTaskAssignedTo, 0, len(tasks))
	for _, task := range tasks {
		randomUserID := ids[rand.Intn(len(ids))]
		// todo one transaction
		err := m.Tasks.AssignTo(ctx, task.ID, randomUserID)
		if err != nil {
			return nil, err
		}

		changeInfo = append(changeInfo, &NewTaskAssignedTo{
			TaskID:        task.PublicID,
			NewAssignedTo: randomUserID,
		})
	}

	return changeInfo, nil
}

func (m *Model) CreateAndAssignTo(ctx context.Context, description, assignTo string) (publicID string, id string, err error) {
	task := &db.Task{
		Description: description,
	}

	// todo one transaction
	publicID, id, err = m.Tasks.Create(ctx, task)
	if err != nil {
		return "", "", err
	}

	err = m.Tasks.AssignTo(ctx, id, assignTo)
	if err != nil {
		return "", "", err
	}

	return publicID, id, nil
}

func (m *Model) MarkAsDone(ctx context.Context, id string) error {
	err := m.Tasks.Completed(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) GetAll(ctx context.Context) ([]*db.Task, error) {
	tasks, err := m.Tasks.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*tasks2.TaskInfo, 0, len(tasks))
	for _, t := range tasks {
		res = append(res, &tasks2.TaskInfo{
			Description: t.Description,
			Done:        t.Done,
		})
	}

	return tasks, nil
}
