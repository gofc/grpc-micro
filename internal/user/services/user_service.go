package services

import (
	"context"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"github.com/pkg/errors"
)

var guestUser = &pb.User{
	Id:      "1",
	Name:    "guest",
	LoginId: "guest",
	Balance: 0,
}

type UserService struct {
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.LoginId == "guest" && req.LoginPass == "1234" {
		return &pb.LoginResponse{
			Token: &pb.Token{
				AccessToken:  "at",
				RefreshToken: "rt",
				ExpiresIn:    1800,
			},
			User: guestUser,
		}, nil
	}
	return nil, errors.New("invalid user loginId or loginPass")
}

func (s *UserService) Me(ctx context.Context, req *pb.MeRequest) (*pb.MeResponse, error) {
	return &pb.MeResponse{User: guestUser}, nil
}

func NewUserService() *UserService {
	return &UserService{}
}
