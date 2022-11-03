package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zarszz/NestAcademy-golang-group-2/server/params"
	"github.com/zarszz/NestAcademy-golang-group-2/server/service"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{
		svc: svc,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	var req params.Pagination
	if err := c.ShouldBindQuery(&req); err != nil {
		WriteInvalidRequestPayloadResponse(c, err.Error())
		return
	}

	resp, err := h.svc.GetProducts(&req)
	if err != nil {
		WriteErrorJsonResponse(c, err)
	}

	WriteJsonResponseGetPaginationSuccess(c, resp)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req params.StoreProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		WriteInvalidRequestPayloadResponse(c, err.Error())
		return
	}

	resp, err := h.svc.CreateProduct(&req)
	if err != nil {
		WriteErrorJsonResponse(c, err)
	}

	WriteJsonResponseSuccess(c, resp)
}

func (h *ProductHandler) FindProductByID(c *gin.Context) {
	productId, isExist := c.Params.Get("productId")
	if !isExist {
		WriteInvalidRequestPayloadResponse(c, "productId is required")
		return
	}
	resp, err := h.svc.FindProductByID(productId)
	if err != nil {
		WriteErrorJsonResponse(c, err)
	}

	WriteJsonResponseGetSuccess(c, resp)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productId := c.Param("productId")
	var req params.StoreProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		WriteInvalidRequestPayloadResponse(c, err.Error())
		return
	}

	resp, err := h.svc.UpdateProduct(&req, productId)
	if err != nil {
		WriteErrorJsonResponse(c, err)
	}

	WriteJsonResponseSuccess(c, resp)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productId, isExist := c.Params.Get("productId")
	if !isExist {
		WriteInvalidRequestPayloadResponse(c, "productId is required")
		return
	}
	resp, err := h.svc.DeleteProduct(productId)
	if err != nil {
		WriteErrorJsonResponse(c, err)
	}

	WriteJsonResponseSuccess(c, resp)
}
