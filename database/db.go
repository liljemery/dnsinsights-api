package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jeremyinoa/dnsinsight-api/configs"
)

var DB *gorm.DB

func Init(cfg *configs.Config) error {
	if DB != nil {
		return nil
	}
	dsn := cfg.PostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
