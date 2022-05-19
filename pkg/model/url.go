package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Uuid         uuid.UUID
	OrgUrl       string `gorm:"not null"`
	ShortUrlPath string `gorm:"not null"`
}
