package request

import (
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