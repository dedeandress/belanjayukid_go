package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       *uuid.UUID `gorm:"Type:uuid;NOT NULL;PRIMARY_KEY;DEFAULT:uuid_generate_v1()" json:"id" db:"id"`
	Email    string     `gorm:"Type:varchar;NULL" json:"email" db:"email"`
	Password string     `gorm:"Type:varchar;NULL" json:"password" db:"password"`
}
