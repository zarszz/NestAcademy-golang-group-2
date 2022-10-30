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

func (t *transactionGormRepository) FindAllByUserID(limit int, page int, userID string) (*[]model.Transaction, *int, error) {
	var transactions []model.Transaction
	offset := (page - 1) * limit
	queryBuilder := t.db.Limit(limit).Offset(offset)
	trx := queryBuilder.Model(model.Transaction{}).Preload("Product").Where("user_id=?", userID).Find(&transactions)
	count := trx.RowsAffected
	err := trx.Error
	if err != nil {
		return nil, nil, err
	}
	countInt := int(count)
	return &transactions, &countInt, nil
}

func (t *transactionGormRepository) FindAll(limit int, page int) (*[]model.Transaction, *int, error) {
	var transactions []model.Transaction
	offset := (page - 1) * limit
	queryBuilder := t.db.Limit(limit).Offset(offset)
	trx := queryBuilder.Model(model.Transaction{}).Preload("Product").Find(&transactions)
	count := trx.RowsAffected
	err := trx.Error
	if err != nil {
		return nil, nil, err
	}
	countInt := int(count)
	return &transactions, &countInt, nil
}

func (t *transactionGormRepository) UpdateStatus(newStatus string, trxID string) error {
	var transaction *model.Transaction
	err := t.db.Where("id=?", trxID).Find(&transaction).Error
	if err != nil {
		return err
	}
	transaction.Status = newStatus
	err = t.db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}
