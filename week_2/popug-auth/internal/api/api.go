package api

import (
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/model"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
)

type server struct {
	mdl *model.Model
	auth.UnimplementedAuthServer
}

// todo implement interface Auth.AuthServer