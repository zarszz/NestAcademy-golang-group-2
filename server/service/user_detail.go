package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zarszz/NestAcademy-golang-group-2/adaptor"
	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"
)

type UserDetailService struct {
	repo              repository.UserDetailRepo
	rajaongkirAdaptor *adaptor.RajaOngkirAdaptor
}

func NewUserDetailService(repo repository.UserDetailRepo, rajaongkirAdaptor *adaptor.RajaOngkirAdaptor) *UserDetailService {
	return &UserDetailService{
		repo:              repo,
		rajaongkirAdaptor: rajaongkirAdaptor,
	}
}

func (u *UserDetailService) CreateUserDetail(user *params.CreateUser, userID string) error {
	cityData, err := u.rajaongkirAdaptor.GetCity(user.CityId, user.ProvinceId)
	if err != nil {
		fmt.Printf("[CreateUserDetail] : error when get data from RajaOngkir : %s", err)
		return custom_error.ErrInternalServer
	}
	userModel := user.ParseToModel(userID)
	userModel.Id = uuid.NewString()
	userModel.CreatedAt = time.Now()
	userModel.UpdatedAt = time.Now()
	userModel.Province = cityData.Rajaongkir.Results.Province
	userModel.City = cityData.Rajaongkir.Results.CityName
	err = u.repo.CreateUserDetail(userModel)
	if err != nil {
		return custom_error.ErrInternalServer
	}
	return nil
}

func (u *UserDetailService) UpdateUser(user *params.CreateUser, userID string) error {
	cityData, err := u.rajaongkirAdaptor.GetCity(user.CityId, user.ProvinceId)
	if err != nil {
		fmt.Printf("[UpdateUser] : error when get data from RajaOngkir : %s", err)
		return custom_error.ErrInternalServer
	}
	userModel := user.ParseToModel(userID)
	userModel.Id = uuid.NewString()
	userModel.UpdatedAt = time.Now()
	userModel.Province = cityData.Rajaongkir.Results.Province
	userModel.City = cityData.Rajaongkir.Results.CityName
	err = u.repo.UpdateUserDetail(userModel, userID)
	if err != nil {
		return custom_error.ErrInternalServer
	}
	return nil
}
