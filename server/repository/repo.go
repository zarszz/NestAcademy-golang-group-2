package repository

import "github.com/zarszz/NestAcademy-golang-group-2/server/model"

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(id string) (*model.User, error)
	FindAllUsers(limit int, page int) (*[]model.User, *int64, error)
	FindUserWithDetailByID(id string) (*model.User, error)
}

type UserDetailRepo interface {
	CreateUserDetail(user *model.UserDetail) error
}
