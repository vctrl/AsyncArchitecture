package api

import (
	"context"
	"github.com/macinnir/jose/crypto"
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/db"
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/model"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
)

type server struct {
	Mdl *model.Model
	auth.UnimplementedAuthServer
}

func New() (*server, error) {
	// todo config
	bts, err := ioutil.ReadFile("/Users/viktor/Repos/async-architecture/week_2/popug-auth/jwtRS256.key")
	if err != nil {
		return nil, err
	}

	rsaPrivate, err := crypto.ParseRSAPrivateKeyFromPEM(bts)
	if err != nil {
		return nil, err
	}

	bts, err = ioutil.ReadFile("/Users/viktor/Repos/async-architecture/week_2/popug-auth/jwtRS256.key.pub")
	if err != nil {
		return nil, err
	}
	rsaPublic, err := crypto.ParseRSAPublicKeyFromPEM(bts)
	if err != nil {
		return nil, err
	}

	sm := model.NewJwtSessionManager(rsaPrivate, rsaPublic)

	if err != nil {
		panic("failed to connect database")
	}

	mdl := &model.Model{
		Sm:    sm,
		Users: db.NewUserRepoSQL("test.db"),
	}

	return &server{Mdl: mdl}, nil
}

func (s *server) CheckSession(ctx context.Context, req *auth.CheckSessionRequest) (*auth.CheckSessionResponse, error) {
	resp, err := s.Mdl.CheckSession(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	return &auth.CheckSessionResponse{
		Status: &auth.Status{
			Code: int32(codes.OK),
			Msg:  "",
		},
		Session: &auth.Session{
			UserId: uint32(resp.UserID),
			Login:  resp.Login,
		},
	}, nil
}

func (s *server) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	pID, id, err := s.Mdl.Register(ctx, req.UserInfo)
	if err != nil {
		return nil, err
	}
	return &auth.RegisterResponse{
		Status: &auth.Status{
			Code: int32(codes.OK),
			Msg:  "",
		},
		PublicId: pID,
		Id:       uint32(id),
	}, nil
}

func (s *server) Login(ctx context.Context, r *auth.LoginRequest) (*auth.LoginResponse, error) {
	token, err := s.Mdl.Login(ctx, r.Login, r.Password)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		Status: &auth.Status{
			Code: int32(codes.OK),
			Msg:  "",
		},
		Token: token,
	}, nil
}

func (s *server) GetUserById(context.Context, *auth.GetUserByIdRequest) (*auth.GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}

func (s *server) UpdateUserById(context.Context, *auth.UpdateUserByIdRequest) (*auth.UpdateUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserById not implemented")
}

func (s *server) DeleteUserById(context.Context, *auth.DeleteUserByIdRequest) (*auth.DeleteUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserById not implemented")
}
