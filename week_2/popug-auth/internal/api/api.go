package api

import (
	"context"
	"github.com/macinnir/jose/crypto"
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/db"
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/model"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
	"google.golang.org/grpc/codes"
	"io/ioutil"
)

type server struct {
	Mdl *model.Model
	auth.UnimplementedAuthServer
}

func New() (*server, error) {
	// todo config and relative path
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
			UserId: resp.UserID,
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
		Id:       id,
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

func (s *server) GetUserById(ctx context.Context, req *auth.GetUserByIdRequest) (*auth.GetUserByIdResponse, error) {
	user, err := s.Mdl.Users.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &auth.GetUserByIdResponse{
		Status: &auth.Status{
			Code: int32(codes.OK),
			Msg:  "",
		},
		UserInfo: &auth.UserInfo{
			Login:    &auth.StringContainer{Value: user.Login},
			Email:    &auth.StringContainer{Value: user.Email},
			FullName: &auth.StringContainer{Value: user.FullName},
			Role:     &auth.StringContainer{Value: user.Role},
		},
	}, nil

}

func (s *server) UpdateUserById(ctx context.Context, req *auth.UpdateUserByIdRequest) (*auth.UpdateUserByIdResponse, error) {
	err := s.Mdl.Users.Update(ctx, req.UserInfo)
	if err != nil {
		return nil, err
	}

	return &auth.UpdateUserByIdResponse{Status: &auth.Status{
		Code: int32(codes.OK),
		Msg:  "",
	}}, nil
}

func (s *server) DeleteUserById(ctx context.Context, req *auth.DeleteUserByIdRequest) (*auth.DeleteUserByIdResponse, error) {
	err := s.Mdl.Users.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &auth.DeleteUserByIdResponse{Status: &auth.Status{
		Code: int32(codes.OK),
		Msg:  "",
	}}, nil
}
