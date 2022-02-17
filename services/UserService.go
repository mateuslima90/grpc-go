package services

import (
	"context"
	"github.com/mateuslima90/grpc-go/pb"
	"github.com/mateuslima90/grpc-go/repository"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (us *UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {

	result, errRepository := us.repository.CreateUser(request.GetUsername(), request.GetName(), request.GetEmail())

	if errRepository != nil {
		return nil, errRepository
	}

	return &pb.User{
		Id:       result.ObjectID,
		Username: result.Username,
		Name:     result.Name,
		Email:    result.Email,
	}, nil
}

func (us *UserService) GetUserByUsername(ctx context.Context, request *pb.User) (*pb.User, error) {
	user, errRepository := us.repository.GetUserByUsername(request.GetUsername())

	if errRepository != nil {
		return nil, errRepository
	}

	return &pb.User{
		Id:       user.ObjectID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (us *UserService) GetUserById(ctx context.Context, request *pb.User) (*pb.User, error) {
	user, errRepository := us.repository.GetUserById(request.Id)

	if errRepository != nil {
		return nil, errRepository
	}

	return &pb.User{
		Id:       user.ObjectID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (us *UserService) GetAllUser(ctx context.Context, request *pb.Empty) (*pb.Users, error) {
	rUsers, errRepository := us.repository.GetAllUser()

	if errRepository != nil {
		return nil, errRepository
	}

	usersList := make([]*pb.User, len(rUsers))
	for i := range rUsers {
		usersList[i] = &pb.User{Id: rUsers[i].ObjectID, Username: rUsers[i].Username,
			Name: rUsers[i].Name, Email: rUsers[i].Email}
	}
	return &pb.Users{Users: usersList}, nil
}
