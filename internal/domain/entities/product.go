package entities

import (
	"errors"
	"product-management/utils"
	"time"
)

type Product struct {
	ID          int
	ManagerID   int
	Category    string
	Price       string
	Name        string
	Description string
	Size        string // small or large
	ExpiredDate time.Time
}

type Metadata struct {
	Cursor int
}

type ProductList struct {
	Metadata
	Items    []Product
}

func (p *Product) validate() error {
	if p.Size != utils.SMALL && p.Size != utils.LARGE {
		return errors.New("잘못된 상품 사이즈 입니다.")
	}
	return nil
}

func NewProduct(id int, managerID int, category string, price string, name string, description string, size string, expiredDate time.Time) *Product {
	return &Product{
		ID: id,
		ManagerID: managerID,
		Category: category,
		Price: price,
		Name: name,
		Description: description,
		Size: size,
		ExpiredDate: expiredDate,
	}
}