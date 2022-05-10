package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	ID        string `gorm:"primaryKey;"`
	Code      string `gorm:"type:varchar(255);not null"`
	Name      string `gorm:"type:varchar(255);not null"`
	Address   string `gorm:"type:varchar(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *Branch) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Branch) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
