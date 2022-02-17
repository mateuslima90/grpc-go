package repository

import "github.com/mateuslima90/grpc-go/entities"

type UserRepository interface {
	CreateUser(username string, name string, email string) (entities.User, error)
	GetUserByUsername(username string) (entities.User, error)
	GetUserById(id string) (entities.User, error)
	GetAllUser() ([]entities.User, error)
}
