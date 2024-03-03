package interfaces

import (
	"product-management/internal/application/command"
	"time"
)

type ManagerService interface {
	SignUp(mng *command.CreateManagerCommand) (statusCode int, err error)
	Login(mng *command.CreateManagerCommand) (token string, tokenExpiration time.Time, statusCode int, err error)
}
