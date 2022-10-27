package gorm_postgres

import (
	"fmt"

	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"
	"gorm.io/gorm"
)

type userDetailGormRepository struct {
	DB *gorm.DB
}

func NewUserDetailGormRepository(DB *gorm.DB) repository.UserDetailRepo {
	return &userDetailGormRepository{
		DB: DB,
	}
}

func (u *userDetailGormRepository) CreateUserDetail(user *model.UserDetail) error {
	err := u.DB.Create(&user).Error
	if err != nil {
		fmt.Printf("[CreateUserDetail] error : %s", err)
		return err
	}
	return nil
}
