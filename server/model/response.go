package model

type Response struct {
	Status         bool   `json:"status"`
	HttpCode       int    `json:"http_code"`
	Message        string `json:"message"`
	ServiceResults ServiceResults
}

type ServiceResults struct {
	ServicesCourier interface{} `json:"services_courier"`
}

type RajaOngkirInquireResponse struct {
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
			CityId     string `json:"city_id"`
			ProvinceId string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"origin_details"`
		DestinationDetails struct {
			CityId     string `json:"city_id"`
			ProvinceId string `json:"province_id"`
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
