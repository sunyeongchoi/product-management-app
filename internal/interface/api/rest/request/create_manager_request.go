package request

type CreateManagerRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
