package service

import (
	"database/sql"
	"log"
	"time"

	"github.com/zarszz/NestAcademy-golang-group-2/helper"
	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"

	"github.com/google/uuid"
)

type UserServices struct {
	repo repository.UserRepo
}

func NewServices(repo repository.UserRepo) *UserServices {
	return &UserServices{
		repo: repo,
	}
}

func (u *UserServices) Register(req *params.Register) error {
	user := req.ParseToModel()

	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "member"

	hash, err := helper.GeneratePassword(user.Password)
	if err != nil {
		log.Printf("[Register] get error when try to generate password %v\n", err)
		return custom_error.ErrInternalServer
	}

	user.Password = hash

	err = u.repo.Register(user)
	if err != nil {
		log.Printf("[Register] get error when save to database %v\n", err)
		return custom_error.ErrInternalServer
	}

	return nil
}

func (u *UserServices) Login(req *params.Login) (*string, error) {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, custom_error.ErrNotFound
	}

	err = helper.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return nil, custom_error.ErrUnauthorized
	}

	token := helper.Token{
		UserId: user.Id,
		Email:  user.Email,
	}

	tokString, err := helper.CreateToken(&token)
	if err != nil {
		return nil, custom_error.ErrInternalServer
	}

	return &tokString, nil
}

func (u *UserServices) FindByID(id string) (*model.User, error) {
	user, err := u.repo.FindUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_error.ErrNotFound
		}
		return nil, custom_error.ErrInternalServer
	}
	return user, nil
}

func (u *UserServices) FindWithDetailByID(id string) (*params.GetUser, error) {
	user, err := u.repo.FindUserWithDetailByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_error.ErrNotFound
		}
		return nil, custom_error.ErrInternalServer
	}
	return makeSingleViewUser(user), nil
}

func (u *UserServices) FindAllUsers(page int, limit int) (*[]params.GetUser, *int64, error) {
	user, count, err := u.repo.FindAllUsers(limit, page)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, custom_error.ErrNotFound
		}
		return nil, nil, custom_error.ErrInternalServer
	}
	return makeListViewUser(user), count, nil
}

func (u *UserServices) FindUserByEmail(email string) (*params.GetUser, error) {
	user, err := u.repo.FindUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_error.ErrNotFound
		}
		return nil, custom_error.ErrInternalServer
	}
	return makeSingleViewUser(user), nil
}

func makeListViewUser(users *[]model.User) *[]params.GetUser {
	var userList []params.GetUser
	for _, user := range *users {
		userList = append(userList, *makeSingleViewUser(&user))
	}
	return &userList
}

func makeSingleViewUser(user *model.User) *params.GetUser {
	return &params.GetUser{
		ID:       user.Id,
		FullName: user.UserDetail.FullName,
		Address: params.UserAddress{
			City: params.LocationIdentity{
				ID:   user.UserDetail.CityId,
				Name: user.UserDetail.City,
			},
			Province: params.LocationIdentity{
				ID:   user.UserDetail.ProvinceId,
				Name: user.UserDetail.Province,
			},
			Street: user.UserDetail.Street,
		},
		Auth: params.UserAuth{
			Email: user.Email,
		},
	}
}
