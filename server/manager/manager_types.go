package manager

type Manager struct {
	ID       int    `json:"id" db:"id"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}
