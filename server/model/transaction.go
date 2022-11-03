package model

import "time"

type InquireTransaction struct {
	Id          string `json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Destination int    `json:"destination"`
	Weight      int    `json:"weight"`
	TotalPrice  int    `json:"total_price"`
	Courier     string `json:"courier"`
}

type RajaOngkirCheckCosts struct {
	Origin      int    `json:"origin"`
	Destination int    `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type InquireTransactionResponse struct {
	Status         int         `json:"status"`
	GeneralInfo    string      `json:"general_info"`
	Message        string      `json:"message"`
	Payload        interface{} `json:"payload"`
	Error          string      `json:"error"`
	AdditionalInfo FailInfo    `json:"additional_info"`
}

type FailInfo struct {
	Message string `json:"message"`
}

//
//var Transactions []InquireTransaction

type Transaction struct {
	Id                    string `json:"id"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Quantity              int     `json:"quantity"`
	Weight                int     `json:"weight"`
	TotalPrice            int     `json:"total_price"`
	DestinationCity       string  `json:"city"`
	DestinationCityID     string  `json:"city_id"`
	DestinationProvince   string  `json:"destination_province"`
	DestinationProvinceID string  `json:"destination_province_id"`
	CourierCode           string  `json:"courier_code"`
	CourierService        string  `json:"courier_service"`
	CourierCost           int     `json:"courier_cost"`
	CourierEstimation     string  `json:"courier_estimation"`
	Status                string  `json:"status"`
	EstimationArrived     string  `json:"estimation_arrived"`
	ProductID             string  `json:"product_id"`
	Product               Product `json:"product"`
	UserID                string  `json:"user_id"`
	User                  User    `json:"user"`
}
