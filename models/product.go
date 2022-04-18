package models

import "github.com/google/uuid"

type Product struct {
	ID uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	SKU string
	Name string
	Stock int
	CategoryID uuid.UUID
	ImageURL string
	ProductDetails []*ProductDetail
	Category Category
}
