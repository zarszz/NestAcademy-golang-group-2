package model

type UserDetail struct {
	BaseModel
	FullName   string `json:"fullname"`
	Gender     string `json:"gender"`
	Contact    string `json:"contact"`
	Street     string `json:"street"`
	CityId     string `json:"city_id"`
	City       string `json:"city"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	UserID     string
}
