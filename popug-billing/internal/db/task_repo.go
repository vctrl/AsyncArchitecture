package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID string `gorm:"primarykey" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`

	PublicID string
	// todo danger: money in int!!!
	// Prize количество денег за выполнение задачи
	Prize int
	// Cost количество денег за назначение задачи
	Cost int
}

type TaskRepo interface {
	GetByID(ctx context.Context, id string) (*Task, error)
	Create(ctx context.Context, task *Task) (string, string, error)
}

func NewTaskRepoSQL(dsn string) TaskRepo {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// todo don't panic
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Task{})

	return &TaskRepositorySQL{db: db}
}

type TaskRepositorySQL struct {
	db *gorm.DB
}

func (r *TaskRepositorySQL) GetByID(ctx context.Context, id string) (*Task, error) {
	var task Task
	// todo it panic if no record found
	db := r.db.WithContext(ctx).First(&task, "public_id = ?", id)
	if db.Error != nil {
		return nil, db.Error
	}

	return &task, nil
}

func (r *TaskRepositorySQL) Create(ctx context.Context, task *Task) (publicID string, id string, err error) {
	task.ID = uuid.NewV4().String()
	db := r.db.WithContext(ctx).Create(task)
	if db.Error != nil {
		return "", "", db.Error
	}

	return task.PublicID, task.ID, nil
}
