package params

import (
	"time"

	"github.com/google/uuid"
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
)

type StoreProductRequest struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Weight      int    `json:"weight"`
	ImageUrl    string `json:"image_url"`
}

func (c *StoreProductRequest) ParseToModel() *model.Product {
	return &model.Product{
		Name:        c.Name,
		Category:    c.Category,
		Description: c.Description,
		Price:       c.Price,
		Stock:       c.Stock,
		Weight:      c.Weight,
		ImageUrl:    c.ImageUrl,
		BaseModel: model.BaseModel{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	ImageUrl    string `json:"image_url"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
