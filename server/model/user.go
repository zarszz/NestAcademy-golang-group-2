package model

type User struct {
	BaseModel
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Users = []User{}
