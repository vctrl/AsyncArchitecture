package model

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/vctrl/async-architecture/schema/events"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vctrl/async-architecture/popug-tasks/internal/db"
	tasks2 "github.com/vctrl/async-architecture/schema/tasks"
)

const (
	createTaskTopic = "task-create-events"
	updateTaskTopic = "task-update-events"
	deleteTaskTopic = "task-delete-events"
	assignTaskTopic = "task-assign-events"
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

	changes := make([]*NewTaskAssignedTo, 0, len(tasks))
	for _, task := range tasks {
		randomUserID := ids[rand.Intn(len(ids))]
		// todo one transaction
		err := m.Tasks.AssignTo(ctx, task.ID, randomUserID)
		if err != nil {
			return nil, err
		}

		changes = append(changes, &NewTaskAssignedTo{
			TaskID:        task.PublicID,
			NewAssignedTo: randomUserID,
		})
	}

	for _, change := range changes {
		evt1 := &events.TaskUpdatedEvent{
			PublicId:   &events.StringContainer{Value: change.TaskID},
			AssignedTo: &events.StringContainer{Value: change.NewAssignedTo},
		}

		evt2 := &events.TaskAssignedEvent{
			TaskPublicId:       change.TaskID,
			AssignedToPublicId: change.NewAssignedTo,
		}

		err = m.produce(evt1, updateTaskTopic)
		if err != nil {
			return nil, err
		}

		err = m.produce(evt2, assignTaskTopic)
		if err != nil {
			return nil, err
		}
	}

	// todo remove
	return nil, nil
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

	msg1 := &events.TaskCreatedEvent{
		PublicId:    publicID,
		AssignedTo:  assignTo,
		Description: description,
		Done:        false,
	}

	err = m.produce(msg1, createTaskTopic)
	if err != nil {
		return "", "", err
	}

	// dirty hack: we have to wait while task will be created before assign to user
	time.Sleep(time.Millisecond * 50)

	err = m.Tasks.AssignTo(ctx, id, assignTo)
	if err != nil {
		return "", "", err
	}

	msg2 := &events.TaskAssignedEvent{
		Meta:               nil,
		TaskPublicId:       publicID,
		AssignedToPublicId: assignTo,
	}

	err = m.produce(msg2, assignTaskTopic)
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

	task, err := m.Tasks.GetByID(ctx, id)
	if err != nil {
		return err
	}

	msg := &events.TaskCompletedEvent{
		Meta:               nil,
		TaskPublicId:       task.PublicID,
		AssignedToPublicId: task.AssignedTo,
	}

	err = m.produce(msg, "task-complete-events")
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

func (m *Model) produce(msg proto.Message, topic string) error {
	evt, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	err = m.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: evt,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
