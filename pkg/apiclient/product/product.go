package product

import (
	"net/http"
	"strconv"

	"product-management/productmgm/common"
	serviceproduct "product-management/productmgm/product"
	"product-management/server/product"

	"github.com/gin-gonic/gin"
)

type apiManager struct{}

func GetProductAPIManager() *apiManager {
	return &apiManager{}
}

func (p apiManager) Register(c *gin.Context) {
	var prod product.Product
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	statusCode, err := serviceproduct.Register(prod)
	if err != nil {
		common.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, common.OKAYMSG, nil).GetProductResponse(c)
}

func (p apiManager) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.NewProductResponse(http.StatusBadRequest, "id를 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
		return
	}
	var updateFields map[string]interface{}
	if err := c.ShouldBindJSON(&updateFields); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	statusCode, err := serviceproduct.Update(id, updateFields)
	if err != nil {
		common.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, common.OKAYMSG, nil).GetProductResponse(c)
}

func (p apiManager) List(c *gin.Context) {
	searchKeyword := c.Query("searchKeyword")
	cursorStr := c.Query("cursor")
	limitStr := c.Query("limit")
	var cursor = 0
	var limit = common.DEFAULT_PAGE_LIMIT
	var err error
	if cursorStr != "" {
		cursor, err = strconv.Atoi(cursorStr)
		if err != nil {
			common.NewProductResponse(http.StatusBadRequest, "cursor를 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
			return
		}
	}
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			common.NewProductResponse(http.StatusBadRequest, "limit을 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
			return
		}
	}
	productList, statusCode, err := serviceproduct.List(searchKeyword, cursor, limit)
	if err != nil {
		common.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, common.OKAYMSG, productList).GetProductResponse(c)
}

func (p apiManager) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.NewProductResponse(http.StatusBadRequest, "id를 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
		return
	}
	prod, statusCode, err := serviceproduct.Get(id)
	if err != nil {
		common.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, common.OKAYMSG, prod).GetProductResponse(c)
}

func (p apiManager) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.NewProductResponse(http.StatusBadRequest, "id를 올바른 타입으로 입력해주세요.", nil).GetProductResponse(c)
		return
	}
	statusCode, err := serviceproduct.Delete(id)
	if err != nil {
		common.NewProductResponse(statusCode, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, common.OKAYMSG, nil).GetProductResponse(c)
}
