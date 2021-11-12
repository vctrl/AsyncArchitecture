package db

import (
	"context"

	"gorm.io/gorm"
)

type Task struct {
	ID          string `gorm:"primarykey" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	AssignedTo  string
	Description string
}

type TaskRepository interface {
	Shuffle(ctx context.Context)
	Create(ctx context.Context, task *Task)
	AssignTo(ctx context.Context, taskID, userID string)

	GetByID(ctx context.Context, id string)
	Update(ctx context.Context, task *Task)
	Delete(ctx context.Context, id string)
}

type TaskRepositorySQL struct {
	db *gorm.DB
}

func (r *TaskRepositorySQL) Shuffle(ctx context.Context) {

}

func (r *TaskRepositorySQL) Create(ctx context.Context, task *Task) {

}

func (r *TaskRepositorySQL) AssignTo(ctx context.Context, taskID, userID string) {

}

func (r *TaskRepositorySQL) GetByID(ctx context.Context, id string) {

}

func (r *TaskRepositorySQL) Update(ctx context.Context, task *Task) {

}

func (r *TaskRepositorySQL) Delete(ctx context.Context, id string) {

}
