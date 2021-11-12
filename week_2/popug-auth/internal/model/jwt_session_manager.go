package model

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Session struct {
	UserID string `json:"user_id"`
	Login  string `json:"login"`
	jwt.StandardClaims
}

type JwtSessionManager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJwtSessionManager(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JwtSessionManager {
	return &JwtSessionManager{
		privateKey: privateKey,
		publicKey:  publicKey,
		accessTTL:  time.Minute,
		//refreshTTL: 0,
	}
}

func (jsm *JwtSessionManager) Create(ctx context.Context, userID string, login string) (string, error) {
	tn := time.Now()
	expiresAt := tn.Add(jsm.accessTTL)
	sess := &Session{
		UserID: userID,
		Login:  login,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  tn.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, sess)
	signed, err := token.SignedString(jsm.privateKey)
	if err != nil {
		return "", err
	}
	return signed, nil
}

func (jsm *JwtSessionManager) Check(ctx context.Context, tokenString string) (*Session, error) {
	payload := &Session{}
	token, err := jwt.ParseWithClaims(tokenString, payload, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok || method.Alg() != "RS256" {
			return nil, fmt.Errorf("bad sign method")
		}
		return jsm.publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return payload, nil
}

func (jsm *JwtSessionManager) Destroy(ctx context.Context, id string) error {
	// ¯\_(ツ)_/¯
	return fmt.Errorf("can't destroy stateless session")
}

func (jsm *JwtSessionManager) DestroyAll(ctx context.Context) error {
	// ¯\_(ツ)_/¯
	return fmt.Errorf("can't destroy stateless sessions")
}
