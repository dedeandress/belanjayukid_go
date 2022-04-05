package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	TotalPrice decimal.Decimal `gorm:"type:decimal(20,2)"`
	Status int
	Date time.Time
}
