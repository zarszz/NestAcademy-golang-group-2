package model

import "time"

type BaseModel struct {
	Id        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
