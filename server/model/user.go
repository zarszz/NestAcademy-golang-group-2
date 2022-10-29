package model

import "time"

type BaseModel struct {
	Id        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type User struct {
	BaseModel
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	UserDetail UserDetail
}

var Users = []User{}
