package command

import "time"

type CreateProductCommand struct {
	ManagerID   int
	Category    string
	Price       string
	Name        string
	Description string
	Size        string // small or large
	ExpiredDate time.Time
}
