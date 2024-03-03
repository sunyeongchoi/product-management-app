package interfaces

import (
	"product-management/internal/application/command"
	"product-management/internal/domain/entities"
)

type ProductService interface {
	Register(prod *command.CreateProductCommand) (statusCode int, err error)
	Update(id int, updateFields map[string]interface{}) (statusCode int, err error)
	List(searchKeyword string, cursor int, limit int) (productList *entities.ProductList, statusCode int, err error)
	Get(id int) (prod *entities.ValidatedProduct, statusCode int, err error)
	Delete(id int) (statusCode int, err error)
}
