package mysql

import "time"

type Manager struct {
	ID       int    `json:"id" db:"id"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}

type Product struct {
	ID          int       `json:"id" db:"id"`
	ManagerID   int       `json:"manager_id" db:"manager_id"`
	Category    string    `json:"category" db:"category"`
	Price       string    `json:"price" db:"price"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Size        string    `json:"size" db:"size"` // small or large
	ExpiredDate time.Time `json:"expired_date" db:"expired_date"`
}
