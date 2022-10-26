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
