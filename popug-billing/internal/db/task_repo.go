package db

import (
	"context"
	"gorm.io/gorm"
)

type Task struct {
	ID string `gorm:"primarykey" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	// todo danger: money in int!!!
	// Prize количество денег за выполнение задачи
	Prize int
	// Cost количество денег за назначение задачи
	Cost int
}

type TaskRepo interface {
	GetByID(ctx context.Context, id string) (*Task, error)
}

type TaskRepositorySQL struct {
	db *gorm.DB
}

func (r *TaskRepositorySQL) GetByID(ctx context.Context, id string) (*Task, error) {
	var task Task
	db := r.db.WithContext(ctx).First(&task, id)
	if db.Error != nil {
		return nil, db.Error
	}

	return &task, nil
}
