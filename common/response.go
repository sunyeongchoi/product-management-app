package common

import (
	"github.com/gin-gonic/gin"

	"product-management/models"
)

// 200 OK Example
// {
// 	"meta":{
// 	  "code": 200, // http status code와 같은 code를 응답으로 전달
// 	  "message":"ok" // 에러 발생시, 필요한 에러 메시지 전달
// 	},
// 	"data":{
// 	  "products":[...]
// 	}
//  }

// 400 Bad Request Example
//  {
// 	"meta":{
// 	  "code": 400,
// 	  "message": "잘못된 상품 사이즈 입니다."
// 	},
// 	"data": null
//  }

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ManagerData struct {
	Managers *models.Manager `json:"managers"`
}

type ProductData struct {
	Products interface{} `json:"products"`
}

type ManagerResponse struct {
	Meta        `json:"meta"`
	ManagerData *ManagerData `json:"data"`
}

type ProductResponse struct {
	Meta        `json:"meta"`
	ProductData *ProductData `json:"data"`
}

func NewManagerResponse(code int, message string, manager *models.Manager) *ManagerResponse {
	if manager == nil {
		return &ManagerResponse{
			Meta: Meta{
				Code:    code,
				Message: message,
			},
			ManagerData: nil,
		}
	}
	return &ManagerResponse{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		ManagerData: &ManagerData{
			Managers: manager,
		},
	}
}

func NewProductResponse(code int, message string, product interface{}) *ProductResponse {
	if product == nil {
		return &ProductResponse{
			Meta: Meta{
				Code:    code,
				Message: message,
			},
			ProductData: nil,
		}
	}
	return &ProductResponse{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		ProductData: &ProductData{
			Products: product,
		},
	}
}

func (m *ManagerResponse) GetManagerResponse(c *gin.Context) {
	c.JSON(m.Meta.Code, m)
}

func (p *ProductResponse) GetProductResponse(c *gin.Context) {
	c.JSON(p.Meta.Code, p)
}
