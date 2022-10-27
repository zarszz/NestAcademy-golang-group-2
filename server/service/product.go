package service

import (
	// "fmt"

	"gorm.io/gorm"

	// "github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"
	"github.com/zarszz/NestAcademy-golang-group-2/server/view"
)

type ProductService struct {
	repo repository.ProductRepo
}

func NewProductServices(repo repository.ProductRepo) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetProducts(req *params.Pagination) (*view.ResponseGetPaginationSuccess, *view.ResponseFailed) {
	products, err := s.repo.GetProducts(req.Limit, req.Offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, view.ErrNotFound("Product not found", err.Error())
		}
		return nil, view.ErrInternalServer("Internal server error", err.Error())
	}

	return view.SuccessGetPagination(view.NewGetAllProductResponse(products), "Success get all products", view.Query{
		Limit: req.Limit,
		Page:  req.Offset,
	}), nil
}

func (s *ProductService) CreateProduct(req *params.StoreProductRequest) (*view.ResponseSuccess, *view.ResponseFailed) {
	err := s.repo.CreateProduct(req.ParseToModel())
	if err != nil {
		return nil, view.ErrInternalServer("Internal server error", err.Error())
	}

	return view.SuccessCreated("Success create product"), nil
}

func (s *ProductService) FindProductByID(id int) (*view.ResponseWithDataSuccess, *view.ResponseFailed) {
	product, err := s.repo.FindProductByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, view.ErrNotFound("Product not found", err.Error())
		}
		return nil, view.ErrInternalServer("Internal server error", err.Error())
	}

	return view.SuccessWithData(view.NewGetProductDetailResponse(product), "Success get product by id"), nil
}

func (s *ProductService) UpdateProduct(req *params.StoreProductRequest) (*view.ResponseSuccess, *view.ResponseFailed) {
	err := s.repo.UpdateProduct(req.ParseToModel())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, view.ErrNotFound("Product not found", err.Error())
		}
		return nil, view.ErrInternalServer("Internal server error", err.Error())
	}

	return view.SuccessCreated("Success update product"), nil
}

func (s *ProductService) DeleteProduct(id int) (*view.ResponseSuccess, *view.ResponseFailed) {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, view.ErrNotFound("Product not found", err.Error())
		}
		return nil, view.ErrInternalServer("Internal server error", err.Error())
	}

	return view.SuccessCreated("Success delete product"), nil
}