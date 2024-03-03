package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management/internal/application/interfaces"
	"product-management/internal/interface/api/rest/request"
	"product-management/internal/interface/api/rest/response"
	middleware2 "product-management/internal/interface/middleware"
	"product-management/utils"
	"strconv"
)

type ProductController struct {
	service interfaces.ProductService
}

func NewProductController(c *gin.Engine, service interfaces.ProductService) *ProductController {
	controller := &ProductController{
		service: service,
	}
	productGroup := c.Group("/management")
	productGroup.Use(middleware2.TokenAuthMiddleware)
	{
		productGroup.POST("/product", controller.Register)
		productGroup.PATCH("/product/:id", controller.Update)
		productGroup.GET("/products", controller.List)
		productGroup.GET("/product/:id", controller.Get)
		productGroup.DELETE("/product/:id", controller.Delete)
	}
	return controller
}

func (pc *ProductController) Register(c *gin.Context) {
	var createProductRequest request.CreateProductRequest
	if err := c.ShouldBindJSON(&createProductRequest); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	statusCode, err := pc.service.Register(&createProductRequest)
	if err != nil {
		response.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	response.NewProductResponse(http.StatusOK, utils.OKAYMSG, nil).GetProductResponse(c)
}

func (pc *ProductController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.NewProductResponse(http.StatusBadRequest, "id를 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
		return
	}
	var updateFields map[string]interface{}
	if err := c.ShouldBindJSON(&updateFields); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	statusCode, err := pc.service.Update(id, updateFields)
	if err != nil {
		response.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	response.NewProductResponse(http.StatusOK, utils.OKAYMSG, nil).GetProductResponse(c)
}

func (pc *ProductController) List(c *gin.Context) {
	searchKeyword := c.Query("searchKeyword")
	cursorStr := c.Query("cursor")
	limitStr := c.Query("limit")
	var cursor = 0
	var limit = utils.DEFAULT_PAGE_LIMIT
	var err error
	if cursorStr != "" {
		cursor, err = strconv.Atoi(cursorStr)
		if err != nil {
			response.NewProductResponse(http.StatusBadRequest, "cursor를 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
			return
		}
	}
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			response.NewProductResponse(http.StatusBadRequest, "limit을 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
			return
		}
	}
	productList, statusCode, err := pc.service.List(searchKeyword, cursor, limit)
	if err != nil {
		response.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	response.NewProductResponse(http.StatusOK, utils.OKAYMSG, productList).GetProductResponse(c)
}

func (pc *ProductController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.NewProductResponse(http.StatusBadRequest, "id를 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
		return
	}
	prod, statusCode, err := pc.service.Get(id)
	if err != nil {
		response.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	response.NewProductResponse(http.StatusOK, utils.OKAYMSG, prod).GetProductResponse(c)
}

func (pc *ProductController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.NewProductResponse(http.StatusBadRequest, "id를 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
		return
	}
	statusCode, err := pc.service.Delete(id)
	if err != nil {
		response.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	response.NewProductResponse(http.StatusOK, utils.OKAYMSG, nil).GetProductResponse(c)
}
