package gorm_postgres

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepoGormPostgres(db *gorm.DB) repository.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetUsers() (*[]model.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *userRepo) Register(user *model.User) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *userRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) FindUserByID(id string) (*model.User, error) {
	var user model.User
	err := u.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) FindUserWithDetailByID(id string) (*model.User, error) {
	var user model.User
	err := u.db.Where("id=?", id).Preload("UserDetail").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) FindAllUsers(limit int, page int) (*[]model.User, *int64, error) {
	var users []model.User
	offset := (page - 1) * limit
	queryBuilder := u.db.Limit(limit).Offset(offset)
	trx := queryBuilder.Model(&model.User{}).Preload("UserDetail").Find(&users)
	count := trx.RowsAffected
	err := trx.Error
	if err != nil {
		return nil, nil, err
	}
	return &users, &count, nil
}
