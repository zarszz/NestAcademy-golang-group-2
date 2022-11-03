package params

import "time"

type Inquire struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Origin      int    `json:"origin"`
	Destination int    `json:"destination"`
	Weight      int    `json:"weight"`
	TotalPrice  int    `json:"total_price"`
	Courier     string `json:"courier"`
}

type InquiryResponse struct {
	Product         InquiryProduct           `json:"product"`
	Quantity        int                      `json:"int"`
	Destination     int                      `json:"destination"`
	Weight          int                      `json:"weight"`
	TotalPrice      int                      `json:"total_price"`
	ServicesCourier []InquiryServicesCourier `json:"services_courier"`
}

type InquiryProduct struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	ImgUrl string `json:"img_url"`
	Price  int    `json:"price"`
}

type InquiryServicesCourier struct {
	Code  string                `json:"code"`
	Name  string                `json:"name"`
	Costs []InquiryServiceCosts `json:"costs"`
}

type InquiryServiceCosts struct {
	Services    string               `json:"services"`
	Description string               `json:"description"`
	Cost        []InquiryServiceCost `json:"cost"`
}

type InquiryServiceCost struct {
	Value      int    `json:"value"`
	Estimation string `json:"estimation"`
	Note       string `json:"note"`
}

type ConfirmTransaction struct {
	ProductID   string                    `json:"product_id"`
	ProductName string                    `json:"product_name"`
	Quantity    int                       `json:"quantity"`
	Destination int                       `json:"destination"`
	Weight      int                       `json:"weight"`
	TotalPrice  int                       `json:"total_price"`
	Courier     ConfirmTransactionCourier `json:"courier"`
}

type ConfirmTransactionCourier struct {
	Code       string `json:"code"`
	Service    string `json:"service"`
	Cost       int    `json:"cost"`
	Estimation string `json:"estimation"`
}

type Transaction struct {
	ID                string      `json:"id"`
	ProductID         string      `json:"product_id"`
	ProductName       string      `json:"product_name"`
	Quantity          string      `json:"quantity"`
	Destination       Destination `json:"destination"`
	Weight            string      `json:"weight"`
	TotalPrice        string      `json:"total_price"`
	Courier           Courier     `json:"courier"`
	Status            string      `json:"status"`
	EstimationArrived string      `json:"estimation_arrived"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

type Destination struct {
	City     string `json:"city"`
	Province string `json:"province"`
}

type Courier struct {
	Code       string `json:"code"`
	Service    string `json:"service"`
	Cost       string `json:"cost"`
	Estimation string `json:"estimation"`
}

type UpdateTrxStatus struct {
	Status string `json:"status"`
}
