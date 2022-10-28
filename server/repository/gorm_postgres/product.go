package gorm_postgres

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepoGormPostgres(db *gorm.DB) repository.ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) GetProducts(limit int, offset int) (*[]model.Product, error) {
	var products []model.Product

	err := p.db.Find(&products).Limit(limit).Offset(offset).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (p *productRepo) CreateProduct(product *model.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepo) FindProductByID(id string) (*model.Product, error) {
	var product model.Product

	err := p.db.Where("id=?", id).Find(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepo) UpdateProduct(product *model.Product) error {
	return p.db.Save(product).Error
}

func (p *productRepo) DeleteProduct(id string) error {
	return p.db.Delete(&model.Product{}, id).Error
}
