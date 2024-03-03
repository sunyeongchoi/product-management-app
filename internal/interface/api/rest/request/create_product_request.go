package request

import (
	"product-management/internal/application/command"
	"time"
)

type CreateProductRequest struct {
	ManagerID   int       `json:"manager_id"`
	Category    string    `json:"category"`
	Price       string    `json:"price"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Size        string    `json:"size"`
	ExpiredDate time.Time `json:"expired_date"`
}

func (req *CreateProductRequest)ToCreateProductCommand() (*command.CreateProductCommand, error) {
	return &command.CreateProductCommand{
		ManagerID: req.ManagerID,
		Category: req.Category,
		Price: req.Price,
		Name: req.Name,
		Description: req.Description,
		Size: req.Size,
		ExpiredDate: req.ExpiredDate,
	}, nil
}