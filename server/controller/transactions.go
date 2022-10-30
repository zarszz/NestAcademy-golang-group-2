package controller

import (
	"net/http"
	"strconv"

	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/service"
	"github.com/zarszz/NestAcademy-golang-group-2/server/view"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	svc service.TransactionServices
}

func NewTransactionController(svc *service.TransactionServices) *TransactionController {
	return &TransactionController{
		svc: *svc,
	}
}

func (t *TransactionController) Inquire(c *gin.Context) {
	var req params.Inquire
	if err := c.ShouldBindJSON(&req); err != nil {
		WriteInvalidRequestPayloadResponse(c, "INQUIRE_FAIL")
		return
	}

	if err := params.Validate(req); err != nil {
		WriteInvalidRequestPayloadResponse(c, "INQUIRE_FAIL")
		return
	}

	inquire := t.svc.Inquire(req)
	if inquire.Status != http.StatusOK {
		views := view.ErrServer(inquire.Status, inquire.AdditionalInfo)
		WriteErrorJsonResponse(c, views)
	} else {
		c.JSON(inquire.Status, inquire)
		//views := view.SuccessWithData(inquire, "INQUIRY_TRANSACTION_SUCCESS")
		//WriteJsonResponseGetSuccess(c, views)
	}
}

func (t *TransactionController) Confirm(c *gin.Context) {
	var confirmation params.ConfirmTransaction
	if err := c.ShouldBindJSON(&confirmation); err != nil {
		WriteInvalidRequestPayloadResponse(c, err.Error())
		return
	}

	userID := c.GetString("USER_ID")
	err := t.svc.ConfirmTransaction(&confirmation, userID)
	if err != nil {
		if err == custom_error.ErrOutOfStock {
			WriteUnprocessableEntityError(c, "stock prodcut not enough", "CONFIRM_TRANSACTION_FAIL")
			return
		}
		payload := view.ErrInternalServer(view.AdditionalInfoError{
			Message: err.Error(),
		}, "CONFIRM_TRANSACTION_FAIL")
		WriteErrorJsonResponse(c, payload)
		return
	}
	payload := view.OperationSuccess("CONFIRM_TRANSACTION_SUCCESS")

	WriteJsonResponseSuccess(c, payload)
}

func (t *TransactionController) FindAllByUserID(c *gin.Context) {
	userID := c.GetString("USER_ID")

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

	trx, count, err := t.svc.GetTransactionsByUserID(limit, page, userID)
	if err != nil {
		payload := view.ErrInternalServer(view.AdditionalInfoError{
			Message: err.Error(),
		}, "INQUIRY_TRANSACTION_FAIL")
		WriteErrorJsonResponse(c, payload)
		return
	}
	query := view.Query{
		Limit: limit,
		Page:  page,
		Total: *count,
	}
	payload := view.SuccessGetPagination(trx, "GET_TRANSACTION_HISTORIES_SUCCESS", query)
	WriteJsonResponseGetPaginationSuccess(c, payload)
}

func (t *TransactionController) FindAll(c *gin.Context) {
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

	trx, count, err := t.svc.GetTransactions(limit, page)
	if err != nil {
		payload := view.ErrInternalServer(view.AdditionalInfoError{
			Message: err.Error(),
		}, "INQUIRY_TRANSACTION_FAIL")
		WriteErrorJsonResponse(c, payload)
		return
	}
	query := view.Query{
		Limit: limit,
		Page:  page,
		Total: *count,
	}
	payload := view.SuccessGetPagination(trx, "GET_TRANSACTION_HISTORIES_SUCCESS", query)
	WriteJsonResponseGetPaginationSuccess(c, payload)
}

func (t *TransactionController) UpdateTrxStatus(c *gin.Context) {
	var update params.UpdateTrxStatus
	if err := c.ShouldBindJSON(&update); err != nil {
		WriteInvalidRequestPayloadResponse(c, err.Error())
		return
	}

	trxID := c.Param("id")
	err := t.svc.UpdateStatus(update.Status, trxID)
	if err != nil {
		payload := view.ErrInternalServer(view.AdditionalInfoError{
			Message: err.Error(),
		}, "UPDATE_STATUS_TRANSACTION_FAIL")
		WriteErrorJsonResponse(c, payload)
		return
	}
	payload := view.OperationSuccess("UPDATE_STATUS_TRANSACTION_SUCCESS")

	WriteJsonResponseSuccess(c, payload)
}
