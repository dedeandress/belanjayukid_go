package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductDetail struct {
	ID uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	Product *Product `gorm:"foreignKey:ProductID;references:ID"`
	ProductUnit *ProductUnit `gorm:"foreignKey:ProductUnitID;references:ID"`
	SellingPrice decimal.Decimal `gorm:"type:decimal(20,8)"`
	PurchasePrice decimal.Decimal `gorm:"type:decimal(20,8)"`
	QuantityPerUnit int
	ProductID uuid.UUID
	ProductUnitID uuid.UUID
}
