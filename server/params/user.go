package params

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
)

type Register struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
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
