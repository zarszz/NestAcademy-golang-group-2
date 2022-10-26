package controller

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/service"
	"github.com/zarszz/NestAcademy-golang-group-2/server/view"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc *service.UserServices
}

func NewUserController(svc *service.UserServices) *UserController {
	return &UserController{
		svc: svc,
	}
}

// func (u *UserController) GetUsers(c *gin.Context) {
// 	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
// 	resp := u.svc.GetUsers()

// 	WriteJsonResponseGin(c, resp)

// }

func (u *UserController) Register(c *gin.Context) {
	var req params.Register
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteInvalidRequestPayloadResponse(c, "REGISTER_FAIL")
		return
	}

	err = params.Validate(req)
	if err != nil {
		WriteInvalidRequestPayloadResponse(c, "REGISTER_FAIL")
		return
	}

	err = u.svc.Register(&req)
	if err != nil {
		info := view.AdditionalInfoError{
			Message: err.Error(),
		}
		payload := view.ErrInternalServer(info, "INTERNAL_SERVER_ERROR")
		WriteErrorJsonResponse(c, payload)
		return
	}

	WriteJsonResponseSuccess(c, view.SuccessCreated("REGISTER_SUCCESS"))
}

func (u *UserController) Login(c *gin.Context) {
	var req params.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		WriteInvalidRequestPayloadResponse(c, "REGISTER_FAIL")
		return
	}

	if err := params.Validate(req); err != nil {
		WriteInvalidRequestPayloadResponse(c, "REGISTER_FAIL")
		return
	}

	token, err := u.svc.Login(&req)
	if err != nil {
		if err == custom_error.ErrInternalServer {
			info := view.AdditionalInfoError{
				Message: "Oopss.. something wrong",
			}
			payload := view.ErrInternalServer(info, "INTERNAL_SERVER_ERROR")
			WriteErrorJsonResponse(c, payload)
			return
		}
		info := view.AdditionalInfoError{
			Message: "invalid username or password",
		}
		payload := view.ErrUnauthorized(info, "UNAUTHORIZED")
		WriteErrorJsonResponse(c, payload)
		return
	}

	payload := map[string]string{"token": *token}
	view := view.SuccessWithData(payload, "LOGIN_SUCCESS")
	WriteJsonResponseGetSuccess(c, view)
}
