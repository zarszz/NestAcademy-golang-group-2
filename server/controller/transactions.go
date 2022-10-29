package controller

import (
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/service"
	"github.com/zarszz/NestAcademy-golang-group-2/server/view"
	"net/http"

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
