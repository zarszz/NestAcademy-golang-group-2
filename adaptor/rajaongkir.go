package adaptor

import (
	"encoding/json"

	"github.com/zarszz/NestAcademy-golang-group-2/config"
	"github.com/zarszz/NestAcademy-golang-group-2/pkg/httpclient"
)

type RajaOngkirGetCity struct {
	Rajaongkir struct {
		Query   interface{} `json:"query"`
		Status  Status      `json:"status"`
		Results Results     `json:"results"`
	} `json:"rajaongkir"`
}

type Query struct {
	Province string `json:"province"`
	ID       string `json:"id"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Results struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type RajaOngkirAdaptor struct {
	httpClient *httpclient.Client
}

func NewRajaOngkirAdapter(baseUrl string) *RajaOngkirAdaptor {
	client := httpclient.NewHttpClient(baseUrl)
	return &RajaOngkirAdaptor{
		httpClient: client,
	}
}

func (r *RajaOngkirAdaptor) GetCity(cityID string, provinceID string) (*RajaOngkirGetCity, error) {
	config, _ := config.LoadConfig(".")
	headers := map[string]string{"key": config.RajaongkirSecret}
	query := map[string]string{"id": cityID, "province": provinceID}
	resp, err := r.httpClient.GetWithHeadersAndQuery("city", &headers, &query)
	if err != nil {
		return nil, err
	}

	var data RajaOngkirGetCity
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
