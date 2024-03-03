package repositories

import "product-management/internal/domain/entities"

type ManagerRepository interface {
	SignUp(manager *entities.ValidatedManager) error
	Get(phone string) (*entities.ValidatedManager, error)
}
