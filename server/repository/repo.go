package repository

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
)

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(id string) (*model.User, error)
	FindAllUsers(limit int, page int) (*[]model.User, *int64, error)
	FindUserWithDetailByID(id string) (*model.User, error)
	FindAllEmployees(page int, limit int) (*[]model.User, *int64, error)
	DeleteByID(id string) error
}

type UserDetailRepo interface {
	CreateUserDetail(user *model.UserDetail) error
	UpdateUserDetail(user *model.UserDetail, userID string) error
	DeleteUserDetailByID(id string) error
}

type TransactionRepo interface {
	Inquire(inquireTransaction *model.InquireTransaction) error
	CekStockProduct(waybill string) ([]model.InquireTransaction, error)
	CekRajaOngkir(inquire params.Inquire) ([]model.InquireTransaction, error)
}

type ProductRepo interface {
	GetProducts(limit int, offset int) (*[]model.Product, error)
	CreateProduct(product *model.Product) error
	FindProductByID(id string) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id string) error
}

type TransactionRepo interface {
	Create(transaction *model.Transaction) error
	FindAllByUserID(limit int, page int, userID string) (*[]model.Transaction, *int, error)
	FindAll(limit int, page int) (*[]model.Transaction, *int, error)
	UpdateStatus(newStatus string, trxID string) error
}
