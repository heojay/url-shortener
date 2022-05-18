package model

import (
	"github.com/google/uuid"
)

type Url struct {
	ID         uuid.UUID
	ShortUrl   string `gorm:"not null"`
	LongUrl    string `gorm:"not null"`
	ClickCount int    `gorm:"default:0"`
}
