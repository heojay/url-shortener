package db

import (
	"fmt"
	"github.com/heojay/url-shortener/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlHost, cfg.MysqlPort, cfg.MysqlDatabase)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
