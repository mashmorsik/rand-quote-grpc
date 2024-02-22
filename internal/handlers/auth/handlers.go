package auth

import (
	"context"
	randquotev1 "github.com/mashmorsik/rand-quote-grpc/api/proto/github.com/mashmorsik/auth-grpc"
	"github.com/mashmorsik/rand-quote-grpc/infrastructure/auth/data"
	"github.com/mashmorsik/rand-quote-grpc/internal/auth"
	"github.com/mashmorsik/rand-quote-grpc/pkg/log"
)

type Server struct {
	user auth.Userer
	randquotev1.AuthServer
}

func NewAuthServer(user auth.Userer) *Server {
	return &Server{user: user}
}

func (as Server) UserAuth(ctx context.Context, req *randquotev1.UserAuthRequest,
) (*randquotev1.UserAuthResponse, error) {
	var responseMsg string

	u := &data.User{
		Id:       0,
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := as.user.UserAuthentication(u)
	if err != nil {
		log.Errf("can't generate token for user: %+v", u)
	}

	if token == "" {
		responseMsg = "User doesn't exist. Create a new user."
	} else {
		responseMsg = token
	}

	return &randquotev1.UserAuthResponse{Token: responseMsg}, nil
}
