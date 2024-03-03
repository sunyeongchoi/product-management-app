package interfaces

import (
	"product-management/internal/domain/entities"
	"product-management/internal/interface/api/rest/request"
)

type ProductService interface {
	Register(prod *request.CreateProductRequest) (statusCode int, err error)
	Update(id int, updateFields map[string]interface{}) (statusCode int, err error)
	List(searchKeyword string, cursor int, limit int) (productList *entities.ProductList, statusCode int, err error)
	Get(id int) (prod *entities.ValidatedProduct, statusCode int, err error)
	Delete(id int) (statusCode int, err error)
}
