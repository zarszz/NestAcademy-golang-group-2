package model

type Product struct {
	BaseModel
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description" gorm:"type:text"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Weight      int    `json:"weight"`
	ImageUrl    string `json:"image_url" gorm:"type:text"`
}
