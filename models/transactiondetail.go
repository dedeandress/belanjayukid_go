package models

import "github.com/google/uuid"

type TransactionDetail struct {
	ID uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	Transaction *Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	ProductDetail *ProductDetail `gorm:"foreignKey:ProductDetailID;references:ID"`
	NumberOfPurchases int
	TransactionID uuid.UUID
	ProductDetailID uuid.UUID
}
