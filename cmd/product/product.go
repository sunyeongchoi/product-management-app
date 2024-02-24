package product

import (
	"net/http"
	"strconv"
	"sync"

	"product-management/common"

	"product-management/models"
	"product-management/sql"
	products "product-management/sql/product"

	"github.com/gin-gonic/gin"
)

type apiManager struct{}

func GetProductAPIManager() *apiManager {
	return &apiManager{}
}

var (
	once          sync.Once
	productDBConn *products.DBProductService
)

func getProductDBConn() *products.DBProductService {
	once.Do(func() {
		productDBConn = products.NewDBProductService(sql.DBConn)
	})
	return productDBConn
}

func (p apiManager) Register(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if product.Size != "small" && product.Size != "large" {
		common.NewProductResponse(http.StatusInternalServerError, "잘못된 상품 사이즈 입니다.", nil).GetProductResponse(c)
		return
	}
	err := getProductDBConn().Register(product)
	if err != nil {
		common.NewProductResponse(http.StatusInternalServerError, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, "ok", nil).GetProductResponse(c)
}

func (p apiManager) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.NewProductResponse(http.StatusInternalServerError, err.Error(), nil).GetProductResponse(c)
		return
	}
	var updateFields map[string]interface{}
	if err := c.ShouldBindJSON(&updateFields); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err = getProductDBConn().Update(id, updateFields)
	if err != nil {
		common.NewProductResponse(http.StatusInternalServerError, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, "ok", nil).GetProductResponse(c)
}

func (p apiManager) List(c *gin.Context) {
	searchKeyword := c.Query("searchKeyword")
	cursorStr := c.Query("cursor")
	limitStr := c.Query("limit")
	var cursor = 0
	var limit = 10 // limit 없으면 기본 10
	var err error
	if cursorStr != "" {
		cursor, err = strconv.Atoi(cursorStr)
		if err != nil {
			common.NewProductResponse(http.StatusInternalServerError, err.Error(), nil).GetProductResponse(c)
			return
		}
	}
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			common.NewProductResponse(http.StatusInternalServerError, err.Error(), nil).GetProductResponse(c)
			return
		}
	}
	productList, err := getProductDBConn().List(searchKeyword, cursor, limit)
	if err != nil {
		common.NewProductResponse(http.StatusInternalServerError, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, "ok", productList).GetProductResponse(c)
}

func (p apiManager) Get(c *gin.Context) {
	id := c.Param("id")
	product, err := getProductDBConn().Get(id)
	if err != nil {
		common.NewProductResponse(http.StatusInternalServerError, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, "ok", product).GetProductResponse(c)
}

func (p apiManager) Delete(c *gin.Context) {
	id := c.Param("id")
	err := getProductDBConn().Delete(id)
	if err != nil {
		common.NewProductResponse(http.StatusInternalServerError, err.Error(), nil).GetProductResponse(c)
		return
	}
	common.NewProductResponse(http.StatusOK, "ok", nil).GetProductResponse(c)
}
