package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

type Config struct {
	MysqlUser     string `envconfig:"MYSQLUSER"`
	MysqlPassword string `envconfig:"MYSQLPASSWORD"`
	MysqlPort     string `envconfig:"MYSQLPORT"`
	MysqlHost     string `envconfig:"MYSQLHOST"`
	MysqlDatabase string `envconfig:"MYSQLDATABASE"`
}

func New() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Debug(err)
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
