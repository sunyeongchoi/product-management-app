package mysql

import "product-management/internal/domain/entities"

func ToDBProduct(product *entities.ValidatedProduct) *Product {
	var p = &Product{
		ID: product.ID,
		ManagerID: product.ManagerID,
		Category: product.Category,
		Price: product.Price,
		Name: product.Name,
		Description: product.Description,
		Size: product.Size,
		ExpiredDate: product.ExpiredDate,
	}
	return p
}

func FromDBProduct(dbProduct *Product) (*entities.ValidatedProduct, error) {
	var p = &entities.Product{
		ID: dbProduct.ID,
		ManagerID: dbProduct.ManagerID,
		Category: dbProduct.Category,
		Price: dbProduct.Price,
		Name: dbProduct.Name,
		Description: dbProduct.Description,
		Size: dbProduct.Size,
		ExpiredDate: dbProduct.ExpiredDate,
	}
	return entities.NewValidatedProduct(p)
}
