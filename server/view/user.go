package view

import "github.com/zarszz/NestAcademy-golang-group-2/server/model"

type UserCreateResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserFindAllResponse struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

func NewUserFindAllResponse(users *[]model.User) []UserFindAllResponse {
	var usersFind []UserFindAllResponse
	for _, user := range *users {
		usersFind = append(usersFind, *parseModelToUserFind(&user))
	}
	return usersFind
}

func parseModelToUserFind(user *model.User) *UserFindAllResponse {
	return &UserFindAllResponse{
		Id:    user.Id,
		Email: user.Email,
	}
}
