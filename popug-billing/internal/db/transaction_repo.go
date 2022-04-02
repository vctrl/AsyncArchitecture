package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Transaction struct {
	ID       string `gorm:"primarykey" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PublicID string
	UserID   string
	TaskID   string
	Amount   int
}

type TransactionRepo interface {
	Create(context.Context, *Transaction) (publicID string, id string, err error)
}

type TransactionRepoSQL struct {
	db *gorm.DB
}

func NewTransactionRepoSQL(dsn string) TransactionRepo {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// todo don't panic
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Transaction{})

	return &TransactionRepoSQL{db: db}

}

func (r *TransactionRepoSQL) Create(ctx context.Context, tran *Transaction) (publicID string, id string, err error) {
	rand.Seed(time.Now().Unix())
	tran.ID = uuid.NewV4().String()
	tran.PublicID = uuid.NewV4().String()

	db := r.db.WithContext(ctx).Create(tran)

	if db.Error != nil {
		return "", "", db.Error
	}

	return tran.PublicID, tran.ID, nil
}
