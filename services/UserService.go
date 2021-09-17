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

	repository.InsertUser2(request.GetName(), request.GetEmail(), request.GetName())

	return &pb.User{
		Id:    "123",
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}

func (*UserService) GetUser(ctx context.Context, request *pb.User) (*pb.User, error) {
	user := repository.GetUser(request.GetName())

	return &pb.User{
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
