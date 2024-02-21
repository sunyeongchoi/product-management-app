package product

import (
	"log"
	"net/http"
	"product-management/models"
	products "product-management/sql/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type apiManager struct{}

func GetProductAPIManager() *apiManager {
	return &apiManager{}
}

func (p apiManager) Register(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println("err: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if product.Size != "small" && product.Size != "large" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "잘못된 상품 사이즈 입니다.",
		})
		return
	}
	err := products.Register(product)
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

func (p apiManager) Update(c *gin.Context) {
	id := c.Param("id")
	idStr, err := strconv.Atoi(id)
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
	err = products.Update(idStr, updateFields)
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
	list, err := products.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    list,
	})
}

func (p apiManager) Get(c *gin.Context) {
	id := c.Param("id")
	product, err := products.Get(id)
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
	err := products.Delete(id)
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

func (p apiManager) Search(c *gin.Context) {
	list, err := products.Search()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    list,
	})
}
