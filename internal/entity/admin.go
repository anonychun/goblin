package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	Id             uuid.UUID
	Name           string
	EmailAddress   string
	PasswordDigest string
	DeletedAt      *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (a *Admin) BeforeCreate(tx *gorm.DB) error {
	a.Id = uuid.Must(uuid.NewV7())
	return nil
}
