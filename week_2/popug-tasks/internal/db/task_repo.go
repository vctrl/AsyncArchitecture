package db

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID          string `gorm:"primarykey" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PublicID    string
	AssignedTo  string
	Description string
	Done        bool
}

type TaskRepo interface {
	Create(ctx context.Context, task *Task) (string, string, error)
	AssignTo(ctx context.Context, taskID, userID string) error
	Completed(ctx context.Context, taskID string) error
	GetAll(ctx context.Context) ([]*Task, error)
	GetByID(ctx context.Context, id string) (*Task, error)
	Update(ctx context.Context, task *Task) error
	DeleteByID(ctx context.Context, id string) error
}

type TaskRepositorySQL struct {
	db *gorm.DB
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

func (r *TaskRepositorySQL) Create(ctx context.Context, task *Task) (publicID string, id string, err error) {
	task.ID = uuid.NewV4().String()
	task.PublicID = uuid.NewV4().String()
	db := r.db.WithContext(ctx).Create(task)
	if db.Error != nil {
		return "", "", db.Error
	}

	return task.PublicID, task.ID, nil
}

func (r *TaskRepositorySQL) AssignTo(ctx context.Context, taskID, userID string) error {
	var task Task
	db := r.db.WithContext(ctx).First(&task, "id = ?", taskID).Update("assigned_to", userID)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (r *TaskRepositorySQL) GetAll(ctx context.Context) ([]*Task, error) {
	var tasks []*Task
	db := r.db.WithContext(ctx).Find(&tasks)
	if db.Error != nil {
		return nil, db.Error
	}

	return tasks, nil
}

func (r *TaskRepositorySQL) Completed(ctx context.Context, taskID string) error {
	db := r.db.WithContext(ctx).First(&Task{}, "id = ?", taskID).Update("done", true)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (r *TaskRepositorySQL) GetByID(ctx context.Context, id string) (*Task, error) {
	var task Task
	db := r.db.WithContext(ctx).First(&task, id)
	if db.Error != nil {
		return nil, db.Error
	}

	return &task, nil
}

func (r *TaskRepositorySQL) Update(ctx context.Context, task *Task) error {
	db := r.db.WithContext(ctx).First(&Task{}, "id = ?", task.ID).Updates(task)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (r *TaskRepositorySQL) DeleteByID(ctx context.Context, id string) error {
	db := r.db.WithContext(ctx).Delete(&Task{}, id)
	if db.Error != nil {
		return db.Error
	}

	return nil
}
