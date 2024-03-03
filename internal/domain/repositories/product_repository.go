package repositories

import "product-management/internal/domain/entities"

type ProductRepository interface {
	Register(product *entities.ValidatedProduct) error
	Update(id int, updateFields map[string]interface{}) error
	List(searchKeyword string, cursor int, limit int) (*entities.ProductList, error)
	Get(id int) (*entities.ValidatedProduct, error)
	Delete(id int) error
}
