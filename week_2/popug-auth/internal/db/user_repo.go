package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicID string
	Login    string
	Password string
	Email    string
	FullName string
	Role     string
}

type UserRepo interface {
	GetByLogin(context context.Context, login string) (*User, error)
	Create(context context.Context, login, password, email, fullName, role string) (publicID string, id uint, err error)
	//Update()
	//Delete()
}

type UserRepoSQL struct {
	DB *gorm.DB
}

func NewUserRepoSQL(dsn string) *UserRepoSQL {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(User{})
	return &UserRepoSQL{DB: db}
}

func (r *UserRepoSQL) GetByLogin(ctx context.Context, login string) (*User, error) {
	var user User
	db := r.DB.First(&user, "login = ?", login)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

func (r *UserRepoSQL) Create(ctx context.Context, login, password, email, fullName, role string) (string, uint, error) {
	uuid := uuid.NewV4()

	user := &User{
		Model:    gorm.Model{},
		PublicID: uuid.String(),
		Login:    login,
		Password: password,
		Email:    email,
		FullName: fullName,
		Role:     role,
	}

	db := r.DB.Create(user)

	if db.Error != nil {
		return "", 0, db.Error
	}

	return uuid.String(), user.ID, nil
}
