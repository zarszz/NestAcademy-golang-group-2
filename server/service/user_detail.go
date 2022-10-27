package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"
)

type UserDetailService struct {
	repo repository.UserDetailRepo
}

func NewUserDetailService(repo repository.UserDetailRepo) *UserDetailService {
	return &UserDetailService{
		repo: repo,
	}
}

func (u *UserDetailService) CreateUserDetail(user *params.CreateUser, userID string) error {
	userModel := user.ParseToModel(userID)
	userModel.Id = uuid.NewString()
	userModel.CreatedAt = time.Now()
	userModel.UpdatedAt = time.Now()
	err := u.repo.CreateUserDetail(userModel)
	if err != nil {
		return custom_error.ErrInternalServer
	}
	return nil
}
