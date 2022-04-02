package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primarykey" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PublicID string
}

type UserRepo interface {
	Create(context context.Context, login, password, email, fullName, role string) (publicID string, id string, err error)
}

type UserRepoSQL struct {
	db *gorm.DB
}

func NewUserRepositorySQL(dsn string) *UserRepoSQL {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to auto migrate")
	}

	return &UserRepoSQL{db: db}
}

func (r *UserRepoSQL) Create(ctx context.Context, login, password, email, fullName, role string) (publicID string, id string, err error) {
	user := &User{
		ID:       uuid.NewV4().String(),
		PublicID: uuid.NewV4().String(),
	}

	db := r.db.WithContext(ctx).Create(user)

	if db.Error != nil {
		return "", "", db.Error
	}

	return user.PublicID, user.ID, nil
}
