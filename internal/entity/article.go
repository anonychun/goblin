package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	Id      uuid.UUID
	Title   string
	Content string
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	a.Id = uuid.Must(uuid.NewV7())
	return nil
}
