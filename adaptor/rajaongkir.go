package adaptor

import (
	"encoding/json"
	"net/url"

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

type RajaongkirCost struct {
	Rajaongkir struct {
		Query struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Weight      int    `json:"weight"`
			Courier     string `json:"courier"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		OriginDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"origin_details"`
		DestinationDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"destination_details"`
		Results []struct {
			Code  string `json:"code"`
			Name  string `json:"name"`
			Costs []struct {
				Service     string `json:"service"`
				Description string `json:"description"`
				Cost        []struct {
					Value int    `json:"value"`
					Etd   string `json:"etd"`
					Note  string `json:"note"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
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
	resp, err := r.httpClient.GetWithHeadersAndQuery("starter/city", &headers, &query)
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

func (r *RajaOngkirAdaptor) CalculateCost(origin string, destination string, weight string, courier string) (*RajaongkirCost, error) {
	config, _ := config.LoadConfig(".")
	headers := map[string]string{"key": config.RajaongkirSecret}
	body := url.Values{}
	body.Add("origin", origin)
	body.Add("destination", destination)
	body.Add("weight", weight)
	body.Add("courier", courier)
	encodedData := body.Encode()
	resp, err := r.httpClient.PostWithHeader("/starter/cost", &headers, &encodedData)
	if err != nil {
		return nil, err
	}

	var data RajaongkirCost
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
