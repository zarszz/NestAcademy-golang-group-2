package repository

import "github.com/zarszz/NestAcademy-golang-group-2/server/model"

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
}

type ProductRepo interface {
	GetProducts(limit int, offset int) (*[]model.Product, error)
	CreateProduct(product *model.Product) error
	FindProductByID(id int) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int) error
}
