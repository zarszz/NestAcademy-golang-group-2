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

func (u *userDetailGormRepository) UpdateUserDetail(newUserData *model.UserDetail, userID string) error {
	var user model.UserDetail
	err := u.DB.Where("user_id=?", userID).Find(&user).Error
	if err != nil {
		fmt.Printf("[CreateUserDetail] error : %s", err)
		return err
	}
	user.City = newUserData.City
	user.CityId = newUserData.CityId
	user.Province = newUserData.Province
	user.ProvinceId = newUserData.ProvinceId
	user.Contact = newUserData.FullName
	user.Gender = newUserData.Gender
	user.Street = newUserData.Street
	user.FullName = newUserData.FullName

	err = u.DB.Save(&user).Error
	if err != nil {
		fmt.Printf("[CreateUserDetail] error : %s", err)
		return err
	}
	return nil
}
