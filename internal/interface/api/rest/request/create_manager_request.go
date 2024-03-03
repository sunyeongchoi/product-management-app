package request

import (
	"product-management/internal/application/command"
)

type CreateManagerRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (req *CreateManagerRequest)ToCreateManagerCommand() (*command.CreateManagerCommand, error) {
	return &command.CreateManagerCommand{
		Phone: req.Phone,
		Password: req.Password,
	}, nil
}