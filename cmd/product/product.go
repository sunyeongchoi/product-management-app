package product

import (
	"net/http"
	"strconv"
	"sync"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "잘못된 상품 사이즈 입니다.",
		})
		return
	}
	err := getProductDBConn().Register(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (p apiManager) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	var updateFields map[string]interface{}
	if err := c.ShouldBindJSON(&updateFields); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err = getProductDBConn().Update(id, updateFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"data":    err.Error(),
			})
			return
		}
	}
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"data":    err.Error(),
			})
			return
		}
	}
	// Like 검색
	productList, err := getProductDBConn().List(searchKeyword, cursor, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	// TODO: 초성 검색

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    productList,
	})
}

func (p apiManager) Get(c *gin.Context) {
	id := c.Param("id")
	product, err := getProductDBConn().Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    product,
	})
}

func (p apiManager) Delete(c *gin.Context) {
	id := c.Param("id")
	err := getProductDBConn().Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
