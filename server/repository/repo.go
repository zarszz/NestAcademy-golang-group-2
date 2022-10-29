package repository

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
)

type UserRepo interface {
	GetUsers() (*[]model.User, error)
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
}

type TransactionRepo interface {
	Inquire(inquireTransaction *model.InquireTransaction) error
	CekStockProduct(waybill string) ([]model.InquireTransaction, error)
	CekRajaOngkir(inquire params.Inquire) ([]model.InquireTransaction, error)
}
