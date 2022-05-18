package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const FileName = "shortener.sqlite"

func New() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(FileName), &gorm.Config{})
}
