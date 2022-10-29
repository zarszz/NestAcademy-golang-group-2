package params

import (
	"errors"
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"

	"github.com/go-playground/validator/v10"
)

type Register struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type CreateUser struct {
	FullName   string `json:"fullname"`
	Gender     string `json:"gender"`
	Contact    string `json:"contact"`
	Street     string `json:"street"`
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
}

type RegisterNewEmployee struct {
	Role       string     `json:"role"`
	Auth       Register   `json:"auth"`
	UserDetail CreateUser `json:"user_detail"`
}

type GetUser struct {
	ID       string      `json:"id"`
	FullName string      `json:"full_name"`
	Address  UserAddress `json:"address"`
	Auth     UserAuth    `json:"auth"`
}

type UserAddress struct {
	City     LocationIdentity `json:"city"`
	Province LocationIdentity `json:"province"`
	Street   string           `json:"street"`
}

type LocationIdentity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserAuth struct {
	Email string `json:"email"`
}

func Validate(u interface{}) error {
	err := validator.New().Struct(u)
	if err == nil {
		return nil
	}
	myErr := err.(validator.ValidationErrors)
	errString := ""
	for _, e := range myErr {
		errString += e.Field() + " is " + e.Tag()
	}
	return errors.New(errString)
}

type UserUpdate struct {
	Email    string
	Password string
}
type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func (u *Register) ParseToModel() *model.User {
	return &model.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *CreateUser) ParseToModel(userID string) *model.UserDetail {
	return &model.UserDetail{
		FullName:   u.FullName,
		Gender:     u.Gender,
		Contact:    u.Contact,
		Street:     u.Street,
		CityId:     u.CityId,
		ProvinceId: u.ProvinceId,
		UserID:     userID,
	}
}
