package mapper

import (
	"product-management/internal/application/command"
	"product-management/internal/domain/entities"
)

func NewManagerResultFromEntity(manager *entities.ValidatedManager) command.ManagerResult {
	return command.ManagerResult{
		ID: manager.ID,
		Phone: manager.Phone,
		Password: manager.Password,
	}
}