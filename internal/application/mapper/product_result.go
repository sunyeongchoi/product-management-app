package mapper

import (
	"product-management/internal/application/command"
	"product-management/internal/domain/entities"
)

func NewProductResultFromEntity(product *entities.ValidatedProduct) command.ProductResult {
	return command.ProductResult{
		ID: product.ID,
		ManagerID: product.ManagerID,
		Category: product.Category,
		Price: product.Price,
		Name: product.Name,
		Description: product.Description,
		Size: product.Size,
		ExpiredDate: product.ExpiredDate,
	}
}
