package models

import "github.com/google/uuid"

type ProductUnit struct {
	ID uuid.UUID `gorm:"NOT NULL;PRIMARY_KEY"`
	Name string
}
