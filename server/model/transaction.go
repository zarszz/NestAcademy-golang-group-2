package model

type Transaction struct {
	BaseModel             BaseModel
	Quantity              int     `json:"quantity"`
	Weight                int     `json:"weight"`
	TotalPrice            int     `json:"total_price"`
	DestinationCity       string  `json:"city"`
	DestinationCityID     string  `json:"city_id"`
	DestinationProvince   string  `json:"destination_province"`
	DestinationProvinceID string  `json:"destination_province_id"`
	CourierCode           string  `json:"courier_code"`
	CourierService        string  `json:"courier_service"`
	CourierCost           string  `json:"courier_cost"`
	CourierEstimation     string  `json:"courier_estimation"`
	Status                string  `json:"status"`
	EstimationArrived     string  `json:"estimation_arrived"`
	Product               Product `json:"product"`
	User                  User    `json:"user"`
}
