package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
	"gorm.io/driver/postgres"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primarykey" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PublicID  string
	CreatedAt time.Time
	UpdatedAt time.Time

	Login    string
	Password string
	Email    string
	FullName string
	Role     string
}

type UserRepo interface {
	GetByLogin(context context.Context, login string) (*User, error)
	Create(context context.Context, login, password, email, fullName, role string) (publicID string, id string, err error)
	GetByID(context context.Context, id string) (*User, error)
	Update(context context.Context, info *auth.UserInfo) error
	Delete(context context.Context, id string) error
}

type UserRepoSQL struct {
	DB *gorm.DB
}

func NewUserRepoSQL(dsn string) *UserRepoSQL {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(User{})
	return &UserRepoSQL{DB: db}
}

func (r *UserRepoSQL) GetByLogin(ctx context.Context, login string) (*User, error) {
	var user User
	db := r.DB.WithContext(ctx).First(&user, "login = ?", login)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

func (r *UserRepoSQL) Create(ctx context.Context, login, password, email, fullName, role string) (publicID string, id string, err error) {
	user := &User{
		ID:        uuid.NewV4().String(),
		PublicID:  uuid.NewV4().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Email:     email,
		FullName:  fullName,
		Role:      role,
	}

	db := r.DB.WithContext(ctx).Create(user)

	if db.Error != nil {
		return "", "", db.Error
	}

	return user.PublicID, user.ID, nil
}

func (r *UserRepoSQL) GetByID(ctx context.Context, id string) (*User, error) {
	var user User
	db := r.DB.WithContext(ctx).First(&user, "id = ?", id)
	if db.Error != nil {
		return nil, db.Error
	}

	return &user, nil
}

func (r *UserRepoSQL) Update(ctx context.Context, info *auth.UserInfo) error {
	db := r.DB.WithContext(ctx).Updates(User{
		UpdatedAt: time.Now(),
		Login:     info.Login.Value,
		Password:  info.Password.Value,
		Email:     info.Email.Value,
		FullName:  info.Email.Value,
		Role:      info.Role.Value,
	})

	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (r *UserRepoSQL) Delete(ctx context.Context, id string) error {
	db := r.DB.WithContext(ctx).Delete(&User{}, id)
	if db.Error != nil {
		return db.Error
	}

	return nil
}
