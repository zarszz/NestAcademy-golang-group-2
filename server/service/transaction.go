package service

import (
	"fmt"
	"strconv"

	"github.com/zarszz/NestAcademy-golang-group-2/adaptor"
	"github.com/zarszz/NestAcademy-golang-group-2/config"
	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"
)

type TransactionService struct {
	userRepo          repository.UserRepo
	productRepo       repository.ProductRepo
	transactionRepo   repository.TransactionRepo
	RajaongkirAdaptor adaptor.RajaOngkirAdaptor
}

func NewTransactionService(rajaongkirAdaptor adaptor.RajaOngkirAdaptor, userRepo repository.UserRepo,
	productRepo repository.ProductRepo, transactionRepo repository.TransactionRepo) *TransactionService {
	return &TransactionService{
		userRepo:          userRepo,
		productRepo:       productRepo,
		transactionRepo:   transactionRepo,
		RajaongkirAdaptor: rajaongkirAdaptor,
	}
}

func (t *TransactionService) InquiryTransaction(inquiry *params.InquiryRequest) (*params.InquiryResponse, error) {
	product, err := t.productRepo.FindProductByID(inquiry.ProductID)
	if err != nil {
		return nil, err
	}

	if product.Stock == 0 {
		return nil, custom_error.ErrOutOfStock
	}

	conf, _ := config.LoadConfig(".")

	totalWeight := product.Weight * inquiry.Quantity

	resp, err := t.RajaongkirAdaptor.CalculateCost(conf.ShopeOriginID, strconv.Itoa(inquiry.Destination), strconv.Itoa(totalWeight), inquiry.Courier)
	if err != nil {
		fmt.Println(err)
	}

	var servicesCourier []params.InquiryServicesCourier

	for _, result := range resp.Rajaongkir.Results {
		var service params.InquiryServicesCourier

		service.Name = result.Name
		service.Code = result.Code

		for _, cost := range result.Costs {
			var serviceCosts params.InquiryServiceCosts
			serviceCosts.Description = cost.Description
			serviceCosts.Services = cost.Service

			for _, detailServiceCost := range cost.Cost {
				var serviceDetailCost params.InquiryServiceCost
				serviceDetailCost.Estimation = detailServiceCost.Etd
				serviceDetailCost.Note = detailServiceCost.Note
				serviceDetailCost.Value = detailServiceCost.Value
				serviceCosts.Cost = append(serviceCosts.Cost, serviceDetailCost)
			}
			service.Costs = append(service.Costs, serviceCosts)
		}
		servicesCourier = append(servicesCourier, service)
	}

	inquiryResult := params.InquiryResponse{
		Product: params.InquiryProduct{
			ID:     product.Id,
			Name:   product.Name,
			ImgUrl: product.ImageUrl,
			Price:  product.Price,
		},
		Quantity:        inquiry.Quantity,
		Destination:     inquiry.Destination,
		Weight:          product.Weight,
		TotalPrice:      product.Price * inquiry.Quantity,
		ServicesCourier: servicesCourier,
	}

	return &inquiryResult, nil
}
