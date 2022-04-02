package model

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/db"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
)

// Model слой бизнес-логики
type Model struct {
	Sm       SessionManager
	Users    db.UserRepo
	Producer *kafka.Producer
}

func (m *Model) Login(ctx context.Context, login, password string) (string, error) {
	user, err := m.Users.GetByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	// todo danger: password in plain text
	if password != user.Password {
		return "", fmt.Errorf("passwords doesn't match")
	}

	token, err := m.Sm.Create(ctx, user.ID, user.Login, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *Model) Register(ctx context.Context, info *auth.UserInfo) (string, string, error) {
	pID, id, err := m.Users.Create(ctx, info.Login.Value, info.Password.Value, info.Email.Value, info.FullName.Value, info.Role.Value)
	if err != nil {
		return "", "", err
	}

	return pID, id, err
}

func (m *Model) CheckSession(ctx context.Context, token string) (*Session, error) {
	sess, err := m.Sm.Check(ctx, token)
	if err != nil {
		return sess, err
	}

	return sess, nil
}
