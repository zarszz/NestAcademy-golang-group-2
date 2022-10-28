package gorm_postgres

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"
	"gorm.io/gorm"
)

type transactionGormRepository struct {
	db *gorm.DB
}

func NewTransactionGormRepository(db *gorm.DB) repository.TransactionRepo {
	return &transactionGormRepository{db: db}
}

func (t *transactionGormRepository) Create(transaction *model.Transaction) error {
	err := t.db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}
