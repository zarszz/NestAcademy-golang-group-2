package model

import "time"

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
