package service

import (
	"bytes"
	"encoding/json"
	"github.com/zarszz/NestAcademy-golang-group-2/config"
	"github.com/zarszz/NestAcademy-golang-group-2/server/model"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"net/http"
)

type TransactionServices struct {
	config config.Config
}

func TransactionServicesNew(config config.Config) *TransactionServices {
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
	url := t.config.RajaOngkirURL + path

	// Generate http body request
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	// Generate http header request
	request = t.generateHeaderRequest(request, t.config.RajaOngkirAPIKey)
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
