package controller

import (
	"strconv"

	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/service"
	"github.com/zarszz/NestAcademy-golang-group-2/server/view"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc           *service.UserServices
	userDetailsvc *service.UserDetailService
}

func NewUserController(svc *service.UserServices, userDetailsvc *service.UserDetailService) *UserController {
	return &UserController{
		svc:           svc,
		userDetailsvc: userDetailsvc,
	}
}

func (u *UserController) GetUsers(c *gin.Context) {
	limitStr, isLimitExist := c.GetQuery("limit")
	pageStr, isPageExist := c.GetQuery("page")

	var limit int
	var page int

	if !isLimitExist {
		limit = 25
	} else {
		limit, _ = strconv.Atoi(limitStr)
	}

	if !isPageExist {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
	}

	users, count, err := u.svc.FindAllUsers(page, limit)

	if err != nil {
		info := view.AdditionalInfoError{
			Message: err.Error(),
		}
		payload := view.ErrInternalServer(info, "INTERNAL_SERVER_ERROR")
		WriteErrorJsonResponse(c, payload)
	}

	payload := view.SuccessWithPaginationData(users, "GET_ALL_USERS_SUCCESS", limit, page, int(*count))
	WriteJsonResponseGetPaginationSuccess(c, payload)
}

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

func (u *UserController) CreateUser(c *gin.Context) {
	var user params.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		WriteInvalidRequestPayloadResponse(c, "CREATED_USER_FAIL")
		return
	}

	userID := c.GetString("USER_ID")
	err := u.userDetailsvc.CreateUserDetail(&user, userID)
	if err != nil {
		if err == custom_error.ErrInternalServer {
			info := view.AdditionalInfoError{
				Message: "Oopss.. something wrong",
			}
			payload := view.ErrInternalServer(info, "INTERNAL_SERVER_ERROR")
			WriteErrorJsonResponse(c, payload)
			return
		}
	}

	view := view.SuccessCreated("CREATED_USER_SUCCESS")
	WriteJsonResponseSuccess(c, view)
}

func (u *UserController) Profile(c *gin.Context) {
	userID := c.GetString("USER_ID")
	user, err := u.svc.FindWithDetailByID(userID)
	if err != nil {
		if err == custom_error.ErrInternalServer {
			info := view.AdditionalInfoError{
				Message: "Oopss.. something wrong",
			}
			payload := view.ErrInternalServer(info, "INTERNAL_SERVER_ERROR")
			WriteErrorJsonResponse(c, payload)
			return
		}
	}
	view := view.SuccessWithData(user, "GET_USER_PROFILE_SUCCESS")
	WriteJsonResponseGetSuccess(c, view)
}

func (u *UserController) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	currentUserEmail := c.GetString("USER_EMAIL")
	user, err := u.svc.FindUserByEmail(email)
	if err != nil {
		if err == custom_error.ErrNotFound {
			info := view.AdditionalInfoError{
				Message: "User not found",
			}
			payload := view.ErrNotFound(info, "GET_USER_BY_EMAIL_FAIL")
			WriteErrorJsonResponse(c, payload)
			return
		}
		info := view.AdditionalInfoError{
			Message: "Oopss.. something wrong",
		}
		payload := view.ErrInternalServer(info, "INTERNAL_SERVER_ERROR")
		WriteErrorJsonResponse(c, payload)
		return
	}

	if email != currentUserEmail {
		info := view.AdditionalInfoError{
			Message: "you dont have access for this resources",
		}
		payload := view.ErrForbidden(info, "GET_USER_BY_EMAIL_FAIL")
		WriteErrorJsonResponse(c, payload)
		return
	}

	payload := view.SuccessWithData(user, "GET_USER_BY_EMAIL_FAIL")
	WriteJsonResponseGetSuccess(c, payload)
}

func (u *UserController) UpdateUserProfile(c *gin.Context) {
	var req params.CreateUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteInvalidRequestPayloadResponse(c, "UPDATE_USER_FAIL")
		return
	}

	err = params.Validate(req)
	if err != nil {
		WriteInvalidRequestPayloadResponse(c, "UPDATE_USER_FAIL")
		return
	}

	userID := c.GetString("USER_ID")
	err = u.userDetailsvc.UpdateUser(&req, userID)
	if err != nil {
		info := view.AdditionalInfoError{
			Message: "Oopss.. something wrong",
		}
		payload := view.ErrInternalServer(info, "INTERNAL_SERVER_ERROR")
		WriteErrorJsonResponse(c, payload)
		return
	}

	payload := view.OperationSuccess("UPDATE_USER_SUCCESS")
	WriteJsonResponseSuccess(c, payload)
}
