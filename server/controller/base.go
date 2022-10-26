package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zarszz/NestAcademy-golang-group-2/server/view"
)

func WriteJsonResponseSuccess(c *gin.Context, payload *view.ResponseSuccess) {
	c.JSON(payload.Status, payload)
}

func WriteJsonResponseGetSuccess(c *gin.Context, payload *view.ResponseWithDataSuccess) {
	c.JSON(payload.Status, payload)
}

func WriteJsonResponseGetPaginationSuccess(c *gin.Context, payload *view.ResponseGetPaginationSuccess) {
	c.JSON(payload.Status, payload)
}

func WriteErrorJsonResponse(c *gin.Context, payload *view.ResponseFailed) {
	c.AbortWithStatusJSON(payload.Status, payload)
}

func WriteInvalidRequestPayloadResponse(c *gin.Context, message string) {
	info := view.InvalidRequestPayload()
	resp := view.ErrBadRequest(info, message)
	WriteErrorJsonResponse(c, resp)
}
