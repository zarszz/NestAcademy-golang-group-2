package view

import (
	"time"

	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
)

type GetProductDetailResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewGetProductDetailResponse(model *model.Product) *GetProductDetailResponse {
	return &GetProductDetailResponse{
		Id:          model.BaseModel.Id,
		Name:        model.Name,
		Category:    model.Category,
		Description: model.Description,
		Price:       model.Price,
		Stock:       model.Stock,
		ImageUrl:    model.ImageUrl,
		CreatedAt:   model.BaseModel.CreatedAt,
		UpdatedAt:   model.BaseModel.UpdatedAt,
	}
}

type GetAllProductResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	ImageUrl    string `json:"image_url"`
}

func NewGetAllProductResponse(products *[]model.Product) *[]GetAllProductResponse {
	var res []GetAllProductResponse

	for _, product := range *products {
		res = append(res, GetAllProductResponse{
			Id:          product.BaseModel.Id,
			Name:        product.Name,
			Category:    product.Category,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			ImageUrl:    product.ImageUrl,
		})
	}

	return &res
}
