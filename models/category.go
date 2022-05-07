package models

import "github.com/google/uuid"

type Category struct {
	ID *uuid.UUID `gorm:"Type:uuid;NOT NULL;PRIMARY_KEY;DEFAULT:uuid_generate_v1()" json:"id" db:"id"`
	Name string
}
