package services

import (
	"context"
	"fmt"

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

	//var deadlineMs = flag.Int("deadline_ms", 20*1000, "Default deadline in milliseconds.")
	//
	//ctx, cancel := context.WithTimeout(ctx, time.Duration(*deadlineMs) * time.Millisecond)
	//defer cancel()

	result := repository.InsertUser(request.GetUsername(), request.GetName(), request.GetEmail())

	return &pb.User{
		Id:       result.ObjectID,
		Username: result.Username,
		Name:     result.Name,
		Email:    result.Email,
	}, nil
}

func (*UserService) GetUserByUsername(ctx context.Context, request *pb.User) (*pb.User, error) {
	user := repository.GetUserByUsername(request.GetUsername())

	return &pb.User{
		Id:       user.ObjectID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (*UserService) GetUserById(ctx context.Context, request *pb.User) (*pb.User, error) {
	//user := repository.GetUser(request.GetName())

	user := repository.GetUserById(request.Id)

	fmt.Println(user)
	return &pb.User{
		Id:       user.ObjectID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}, nil
}

func (*UserService) GetAllUser(ctx context.Context, request *pb.Empty) (*pb.Users, error) {
	rUsers := repository.GetAllUser()
	//fmt.Println(users)
	usersList := make([]*pb.User, len(rUsers))
	for i := range rUsers {
		//fmt.Println(users[i])
		usersList[i] = &pb.User{Id: rUsers[i].ObjectID, Username: rUsers[i].Username, Name: rUsers[i].Name, Email: rUsers[i].Email}
	}
	return &pb.Users{
		Users: usersList,
	}, nil
}
