package command

import "time"

type ProductResult struct {
	ID          int
	ManagerID   int
	Category    string
	Price       string
	Name        string
	Description string
	Size        string // small or large
	ExpiredDate time.Time
}
