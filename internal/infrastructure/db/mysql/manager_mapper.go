package mysql

import "product-management/internal/domain/entities"

func ToDBManager(product *entities.ValidatedManager) *Manager {
	var m = &Manager{
		ID: product.ID,
		Phone: product.Phone,
		Password: product.Password,
	}
	return m
}

func FromDBManager(dbManager *Manager) (*entities.ValidatedManager, error) {
	var m = &entities.Manager{
		ID: dbManager.ID,
		Phone: dbManager.Phone,
		Password: dbManager.Password,
	}
	return entities.NewValidatedManager(m)
}
