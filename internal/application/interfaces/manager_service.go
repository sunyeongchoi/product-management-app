package interfaces

import (
	"product-management/internal/interface/api/rest/request"
	"time"
)

type ManagerService interface {
	SignUp(mng *request.CreateManagerRequest) (statusCode int, err error)
	Login(mng *request.CreateManagerRequest) (token string, tokenExpiration time.Time, statusCode int, err error)
}
