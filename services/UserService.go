package services

import (
	"context"

	"github.com/mateuslima90/grpc-go/pb"
	"github.com/mateuslima90/grpc-go/repository"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {

	repository.InsertUser(request.GetName(), request.GetEmail())

	return &pb.User{
		Id:    "123",
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}
