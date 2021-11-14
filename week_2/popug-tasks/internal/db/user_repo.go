package db

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var user User
var users []User

type User struct {
	ID       string `gorm:"primarykey" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PublicID string

	Login    string
	Email    string
	FullName string
}

type UserRepo interface {
	GetAllUserIDs(ctx context.Context) ([]string, error)
	IsExist(ctx context.Context, id string) (bool, error)
}

type UserRepositorySQL struct {
	db *gorm.DB
}

func NewUserRepositorySQL(dsn string) *UserRepositorySQL {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to auto migrate")
	}

	return &UserRepositorySQL{db: db}
}

func (r *UserRepositorySQL) GetAllUserIDs(ctx context.Context) ([]string, error) {
	db := r.db.WithContext(ctx).Select("id").Find(&users)
	if db.Error != nil {
		return nil, db.Error
	}

	ids := make([]string, 0, db.RowsAffected)
	rows, err := db.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u User
		err = r.db.WithContext(ctx).ScanRows(rows, &u)
		if err != nil {
			return nil, err
		}
		ids = append(ids, u.ID)
	}

	return ids, nil
}

func (r *UserRepositorySQL) IsExist(ctx context.Context, id string) (bool, error) {
	db := r.db.WithContext(ctx).First(id, &user)
	if db.Error != nil {
		return false, db.Error
	}

	if db.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
