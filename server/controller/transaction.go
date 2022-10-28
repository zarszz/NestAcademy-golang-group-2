package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zarszz/NestAcademy-golang-group-2/server/custom_error"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/service"
	"github.com/zarszz/NestAcademy-golang-group-2/server/view"
)

type TransactionController struct {
	svc service.TransactionService
}

func NewTransactionController(svc *service.TransactionService) *TransactionController {
	return &TransactionController{svc: *svc}
}

func (t *TransactionController) Inquiry(c *gin.Context) {
	var inquiry params.InquiryRequest
	if err := c.ShouldBindJSON(&inquiry); err != nil {
		WriteInvalidRequestPayloadResponse(c, err.Error())
		return
	}
	res, err := t.svc.InquiryTransaction(&inquiry)
	if err != nil {
		if err == custom_error.ErrOutOfStock {
			WriteUnprocessableEntityError(c, "stock prodcut not enough", "INQUIRY_TRANSACTION_FAIL")
			return
		}
		payload := view.ErrInternalServer(view.AdditionalInfoError{
			Message: err.Error(),
		}, "INQUIRY_TRANSACTION_FAIL")
		WriteErrorJsonResponse(c, payload)
		return
	}
	payload := view.SuccessWithData(res, "INQUIRY_TRANSACTION_SUCCESS")

	WriteJsonResponseGetSuccess(c, payload)
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
