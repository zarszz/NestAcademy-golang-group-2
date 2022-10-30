package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zarszz/NestAcademy-golang-group-2/adaptor"
	"github.com/zarszz/NestAcademy-golang-group-2/config"
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/repository"
)

type TransactionServices struct {
	config            config.Config
	userRepo          repository.UserRepo
	productRepo       repository.ProductRepo
	transactionRepo   repository.TransactionRepo
	RajaongkirAdaptor adaptor.RajaOngkirAdaptor
}

func TransactionServicesNew(config config.Config, rajaongkirAdaptor adaptor.RajaOngkirAdaptor,
	userRepo repository.UserRepo, productRepo repository.ProductRepo, transactionRepo repository.TransactionRepo) *TransactionServices {
	return &TransactionServices{
		config: config,
	}
}

const (
	ProductCodeRegular = "REG"
	PathServiceRates   = "/serviceRates"
	PathCheckOngkir    = "/cost"
	PathRequestPickup  = "/requestPickup"
	DateFormat         = "20060102"
	TrackingCodeClosed = 250
	TrackingCodeRetur  = 255
)

func (t *TransactionServices) Inquire(req params.Inquire) model.InquireTransactionResponse {
	// set default result
	var result = model.InquireTransactionResponse{
		Status:      http.StatusOK,
		GeneralInfo: "NooBee-Shop",
		Message:     "INQUIRY_TRANSACTION_SUCCESS",
	}

	// melakukan pengecekan untuk estimasi ongkos kirim
	cekOngkir := t.CekRajaOngkir(req)
	if cekOngkir.HttpCode != http.StatusOK {
		result.Status = cekOngkir.HttpCode
		result.Message = "INQUIRY_TRANSACTION_FAIL"
		result.Error = "INQUIRY_TRANSACTION_FAIL"
		result.AdditionalInfo.Message = cekOngkir.Message

		return result
	}
	result.Payload = cekOngkir.ServiceResults

	return result
}

func (t *TransactionServices) CekRajaOngkir(bodyRequest params.Inquire) model.Response {
	// set default result
	var result = model.Response{
		Status:   true,
		HttpCode: http.StatusOK,
		Message:  "Success",
	}

	// check ongkir to raja ongkir
	var requestRajaOngkir = model.RajaOngkirCheckCosts{
		Origin:      bodyRequest.Origin,
		Destination: bodyRequest.Destination,
		Weight:      bodyRequest.Weight,
		Courier:     bodyRequest.Courier,
	}
	// convert body request to json
	request, _ := json.Marshal(requestRajaOngkir)

	// Integrate to raja ongkir
	var dataResponse model.RajaOngkirInquireResponse
	response, err := t.integrateToRajaOngkir(PathCheckOngkir, request)

	// check integration result
	if err == nil {
		// Decode json response to struct
		_ = json.NewDecoder(response.Body).Decode(&dataResponse)

		// check result
		if response.StatusCode == http.StatusOK {
			// find product code regular index on slice
			//slices
			//index := slices.IndexFunc(dataResponse.Rajaongkir.Results, func(s model.RajaOngkirInquireResponse) bool { return s.ProductCode == ProductCodeRegular })

			// set result data
			result.ServiceResults.ServicesCourier = dataResponse.Rajaongkir.Results
		} else {
			// set error response
			result.HttpCode = response.StatusCode
			result.Message = dataResponse.Rajaongkir.Status.Description
		}
	} else {
		// set integration error response
		result.Status = false
		result.HttpCode = http.StatusInternalServerError
		result.Message = "Internal Server Error. Please contact Administrator!"
	}

	// return result
	return result
}

func (t *TransactionServices) integrateToRajaOngkir(path string, bodyRequest []byte) (*http.Response, error) {
	var client = &http.Client{}

	// Set request url
	url := t.config.RajaongkirBaseUrl + path

	// Generate http body request
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	// Generate http header request
	request = t.generateHeaderRequest(request, t.config.RajaongkirSecret)
	// Request to raja ongkir
	response, err := client.Do(request)

	// return result
	return response, err
}

func (t *TransactionServices) generateHeaderRequest(request *http.Request, apiKey string) *http.Request {
	request.Header.Set("content-Type", "application/json")
	request.Header.Set("key", apiKey)

	return request
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

func (t *TransactionService) GetTransactionsByUserID(limit int, page int, userID string) (*[]params.Transaction, *int, error) {
	trx, count, err := t.transactionRepo.FindAllByUserID(limit, page, userID)
	if err != nil {
		fmt.Printf("[GetTransactionsByUserID] error : %v", err)
		return nil, nil, err
	}
	return makeListTransactionView(trx), count, nil
}

func (t *TransactionService) GetTransactions(limit int, page int) (*[]params.Transaction, *int, error) {
	trx, count, err := t.transactionRepo.FindAll(limit, page)
	if err != nil {
		fmt.Printf("[GetTransactions] error : %v", err)
		return nil, nil, err
	}
	return makeListTransactionView(trx), count, nil
}

func (t *TransactionService) UpdateStatus(newStatus string, trxID string) error {
	err := t.transactionRepo.UpdateStatus(newStatus, trxID)
	if err != nil {
		fmt.Printf("[UpdateStatus] error : %v", err)
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

func makeListTransactionView(trxs *[]model.Transaction) *[]params.Transaction {
	var res []params.Transaction
	for _, trx := range *trxs {
		res = append(res, *makeSingleTransactionView(&trx))
	}
	return &res
}

func makeSingleTransactionView(trx *model.Transaction) *params.Transaction {
	return &params.Transaction{
		ID:          trx.Id,
		ProductID:   trx.ProductID,
		ProductName: trx.Product.Name,
		Quantity:    strconv.Itoa(trx.Quantity),
		Destination: params.Destination{
			City:     trx.DestinationCity,
			Province: trx.DestinationProvince,
		},
		Weight:     strconv.Itoa(trx.Weight),
		TotalPrice: strconv.Itoa(trx.TotalPrice),
		Courier: params.Courier{
			Code:       trx.CourierCode,
			Service:    trx.CourierService,
			Cost:       strconv.Itoa(trx.CourierCost),
			Estimation: trx.CourierEstimation,
		},
		Status:            trx.Status,
		EstimationArrived: trx.EstimationArrived,
		CreatedAt:         trx.CreatedAt,
		UpdatedAt:         trx.UpdatedAt,
	}
}
