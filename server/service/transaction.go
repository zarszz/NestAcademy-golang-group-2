package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zarszz/NestAcademy-golang-group-2/adaptor"
	"github.com/zarszz/NestAcademy-golang-group-2/config"
	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
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
		ServicesCourier: *parseToStruct(resp),
	}

	return &inquiryResult, nil
}

func (t *TransactionService) ConfirmTransaction(confirmTransaction *params.ConfirmTransaction, userID string) error {
	product, err := t.productRepo.FindProductByID(confirmTransaction.ProductID)
	if err != nil {
		return err
	}

	user, err := t.userRepo.FindUserByID(userID)
	if err != nil {
		return err
	}

	if product.Stock == 0 {
		return custom_error.ErrOutOfStock
	}

	conf, _ := config.LoadConfig(".")

	totalWeight := confirmTransaction.Quantity * product.Weight

	resp, err := t.RajaongkirAdaptor.CalculateCost(conf.ShopeOriginID, strconv.Itoa(confirmTransaction.Destination), strconv.Itoa(totalWeight), confirmTransaction.Courier.Code)
	if err != nil {
		fmt.Println(err)
	}

	courierServices := *parseToStruct(resp)

	var selectedService params.InquiryServiceCosts
	for _, courier := range courierServices {

		// select specific courier (such as jne, jnt, etc)
		if courier.Code == confirmTransaction.Courier.Code {
			for _, service := range courier.Costs {

				// select specific service in courier (such as REG, YES, etc)
				if service.Services == confirmTransaction.Courier.Service {
					selectedService = service
				}
			}
		}
	}

	totalPrice := (confirmTransaction.Quantity * product.Price) + selectedService.Cost[0].Value

	var estimationArrived string
	var estimationDay string
	estimation := selectedService.Cost[0].Estimation
	if strings.Contains(estimation, "-") {
		splitted := strings.Split(estimation, "-")
		bestInt, _ := strconv.Atoi(splitted[0])
		best := time.Now().AddDate(0, 0, bestInt)
		worstInt, _ := strconv.Atoi(splitted[1])
		worst := time.Now().AddDate(0, 0, worstInt)
		estimationArrived = fmt.Sprintf("%s - %s", best.Format(time.RFC3339), worst.Format(time.RFC3339))
		estimationDay = fmt.Sprintf("%s days", estimationDay)
	} else {
		estimationArrived = time.Now().Format(time.RFC3339)
		estimationDay = "1 day"
	}

	transaction := model.Transaction{
		Id:                    uuid.NewString(),
		CreatedAt:             time.Now(),
		Quantity:              confirmTransaction.Quantity,
		Weight:                totalWeight,
		TotalPrice:            totalPrice,
		DestinationCity:       resp.Rajaongkir.DestinationDetails.CityName,
		DestinationCityID:     resp.Rajaongkir.DestinationDetails.CityID,
		DestinationProvince:   resp.Rajaongkir.DestinationDetails.Province,
		DestinationProvinceID: resp.Rajaongkir.DestinationDetails.ProvinceID,
		CourierCode:           confirmTransaction.Courier.Code,
		CourierService:        selectedService.Services,
		CourierCost:           selectedService.Cost[0].Value,
		CourierEstimation:     estimationDay,
		Status:                "WAITING",
		EstimationArrived:     estimationArrived,
		ProductID:             product.Id,
		Product:               *product,
		UserID:                user.BaseModel.Id,
		User:                  *user,
	}

	err = t.transactionRepo.Create(&transaction)
	if err != nil {
		return err
	}

	return nil
}

func parseToStruct(rajaongkirResp *adaptor.RajaongkirCost) *[]params.InquiryServicesCourier {
	var servicesCourier []params.InquiryServicesCourier

	for _, result := range rajaongkirResp.Rajaongkir.Results {
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
	return &servicesCourier
}
