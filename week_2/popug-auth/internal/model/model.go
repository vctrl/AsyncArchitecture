package model

import (
	"context"
	"fmt"
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/db"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
)

// Model слой бизнес-логики
type Model struct {
	Sm    SessionManager
	Users db.UserRepo
}

func (m *Model) Login(ctx context.Context, login, password string) (string, error) {
	user, err := m.Users.GetByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	if password != user.Password {
		return "", fmt.Errorf("passwords doesn't match")
	}

	token, err := m.Sm.Create(ctx, user.ID, user.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *Model) Register(ctx context.Context, info *auth.UserInfo) (string, uint, error) {
	pID, id, err := m.Users.Create(ctx, info.Login.Value, info.Password.Value, info.Email.Value, info.FullName.Value, info.Role.Value)
	if err != nil {
		return "", 0, err
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
